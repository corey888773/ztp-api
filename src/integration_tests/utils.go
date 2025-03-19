package integration_tests

import (
	"context"
	"log"
	"testing"

	"github.com/corey888773/ztp-api/src/data"
	"github.com/corey888773/ztp-api/src/util"
	"github.com/stretchr/testify/assert"
	"go.mongodb.org/mongo-driver/bson"
)

func ClearDb(t *testing.T) {
	config, err := util.LoadConfig(".")
	if err != nil {
		assert.NoError(t, err)
	}

	ctx := context.Background()
	client, err := data.InitMongoDB(ctx, config)
	assert.NoError(t, err)

	dbNames, err := client.ListDatabaseNames(ctx, bson.M{})
	if err != nil {
		assert.NoError(t, err)
	}

	for _, dbName := range dbNames {
		if dbName == "admin" || dbName == "local" || dbName == "config" {
			continue
		}
		if err := client.Database(dbName).Drop(ctx); err != nil {
			assert.NoError(t, err)
		}
		log.Printf("Dropped database: %s", dbName)
	}
}
