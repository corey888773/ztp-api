package data

import (
	"context"
	"time"

	"github.com/corey888773/ztp-api/src/util"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func InitMongoDB(appCtx context.Context, config util.Config) (*mongo.Client, error) {
	// Connect to MongoDB
	mongoCtx, _ := context.WithTimeout(appCtx, 10*time.Second)
	return mongo.Connect(mongoCtx, options.Client().ApplyURI(config.MongoUri))
}
