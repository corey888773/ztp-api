package data

// import (
//
//	"time"
//
//	"github.com/corey888773/ztp-api/src/util"
//	"go.mongodb.org/mongo-driver/bson"
//	"go.mongodb.org/mongo-driver/mongo"
//
// )
//type ProductChange struct {
//	ID         string   `bson:"_id,omitempty"`
//	ProductId  string   `bson:"product_id"`
//	Change     string   `bson:"change"`
//	State      *Product `bson:"state"`
//	CreatedAt  int64    `bson:"created_at"`
//	ModifiedBy string   `bson:"modified_by"`
//}

//
//type HistoryRepository interface {
//	GetProductHistory(id string) ([]ProductChange, error)
//	RegisterProductChange(change ProductChange) error
//}
//
//type historyRepository struct {
//	historyCollection *mongo.Collection
//}
//
//func (p *historyRepository) RegisterProductChange(change ProductChange) error {
//	_, err = p.historyCollection.InsertOne(sessCtx)
//	if err != nil {
//		return nil, err
//	}
//}
//
//func (p *historyRepository) GetProductHistory(id string) ([]ProductChange, error) {
//	ctx, cancel := util.CreateContext()
//	defer cancel()
//
//	history := []ProductChange{}
//	cursor, err := p.historyCollection.Find(ctx, bson.M{"product_id": id})
//	if err != nil {
//		return nil, err
//	}
//
//	if err = cursor.All(ctx, &history); err != nil {
//		return nil, err
//	}
//
//	return history, nil
//}
//
//func NewHistoryRepository(historyCollection *mongo.Collection) (HistoryRepository, error) {
//	return &historyRepository{
//		historyCollection: historyCollection,
//	}, nil
//}
