package data

import (
	"fmt"
	"time"

	"github.com/corey888773/ztp-api/src/util"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type ProductChange struct {
	ID         string   `bson:"_id,omitempty"`
	ProductId  string   `bson:"product_id"`
	Change     string   `bson:"change"`
	State      *Product `bson:"state"`
	CreatedAt  int64    `bson:"created_at"`
	ModifiedBy string   `bson:"modified_by"`
}

type Product struct {
	ID       string  `bson:"_id,omitempty" json:"id"`
	Category string  `bson:"category" json:"category"`
	Name     string  `bson:"name" json:"name"`
	Price    float64 `bson:"price" json:"price"`
	Quantity int     `bson:"quantity" json:"quantity"`
}

type ProductRepository interface {
	FindAll() ([]Product, error)
	FindById(id string) (*Product, error)
	Create(product Product) error
	Update(product Product) error
	Delete(id string) error
	GetProductHistory(id string) ([]ProductChange, error)
}

type productRepository struct {
	client            *mongo.Client
	productCollection *mongo.Collection
	changeCollection  *mongo.Collection
}

func (p *productRepository) GetProductHistory(id string) ([]ProductChange, error) {
	ctx, cancel := util.CreateContext()
	defer cancel()

	history := []ProductChange{}
	cursor, err := p.changeCollection.Find(ctx, bson.M{"product_id": id})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &history); err != nil {
		return nil, err
	}

	return history, nil
}

func (p *productRepository) FindAll() ([]Product, error) {
	ctx, cancel := util.CreateContext()
	defer cancel()

	trips := []Product{}
	cursor, err := p.productCollection.Find(ctx, bson.M{})
	if err != nil {
		return nil, err
	}

	if err = cursor.All(ctx, &trips); err != nil {
		return nil, err
	}

	return trips, nil
}

func (p *productRepository) FindById(id string) (*Product, error) {
	ctx, cancel := util.CreateContext()
	defer cancel()

	var product Product
	err := p.productCollection.FindOne(ctx, bson.M{"_id": id}).Decode(&product)
	if err != nil {
		return nil, err
	}

	return &product, nil
}

func (p *productRepository) Create(product Product) error {
	ctx, cancel := util.CreateContext()
	defer cancel()

	session, err := p.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		product.ID = primitive.NewObjectID().Hex()
		_, err := p.productCollection.InsertOne(sessCtx, product)
		if err != nil {
			return nil, err
		}

		productChange := ProductChange{
			ProductId:  product.ID,
			Change:     "created",
			State:      &product,
			CreatedAt:  time.Now().Unix(),
			ModifiedBy: "system",
		}
		_, err = p.changeCollection.InsertOne(sessCtx, productChange)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	return err
}

func (p *productRepository) Update(product Product) error {
	ctx, cancel := util.CreateContext()
	defer cancel()

	session, err := p.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := p.productCollection.ReplaceOne(sessCtx, bson.M{"_id": product.ID}, product)
		if err != nil {
			return nil, err
		}

		productChange := ProductChange{
			ProductId:  product.ID,
			Change:     "updated",
			State:      &product,
			CreatedAt:  time.Now().Unix(),
			ModifiedBy: "system",
		}
		_, err = p.changeCollection.InsertOne(sessCtx, productChange)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	return err
}

func (p *productRepository) Delete(id string) error {
	ctx, cancel := util.CreateContext()
	defer cancel()

	session, err := p.client.StartSession()
	if err != nil {
		return err
	}
	defer session.EndSession(ctx)

	callback := func(sessCtx mongo.SessionContext) (interface{}, error) {
		_, err := p.productCollection.DeleteOne(sessCtx, bson.M{"_id": id})
		if err != nil {
			return nil, err
		}

		productChange := ProductChange{
			ProductId:  id,
			Change:     "deleted",
			CreatedAt:  time.Now().Unix(),
			ModifiedBy: "system",
			State:      nil,
		}
		_, err = p.changeCollection.InsertOne(sessCtx, productChange)
		if err != nil {
			return nil, err
		}

		return nil, nil
	}

	_, err = session.WithTransaction(ctx, callback)
	return err
}

func NewProductRepository(client *mongo.Client, productsCollection, changeCollection *mongo.Collection) (ProductRepository, error) {
	ctx, cancel := util.CreateContext()
	defer cancel()

	// Create unique index on "name" for products.
	_, err := productsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
		Keys:    bson.M{"name": 1},
		Options: options.Index().SetUnique(true),
	})
	if err != nil {
		return nil, fmt.Errorf("failed to create name index: %w", err)
	}

	return &productRepository{
		client:            client,
		productCollection: productsCollection,
		changeCollection:  changeCollection,
	}, nil
}
