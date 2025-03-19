package app

import (
	"context"
	"reflect"
	"strings"

	"github.com/corey888773/ztp-api/src/api"
	"github.com/corey888773/ztp-api/src/data"
	"github.com/corey888773/ztp-api/src/types"
	"github.com/corey888773/ztp-api/src/util"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
)

func CreateApp(appCtx context.Context) (*api.Srv, error) {
	config, err := util.LoadConfig(".")
	if err != nil {
		return nil, err
	}

	// init mongoDB client
	mongoClient, err := data.InitMongoDB(appCtx, config)

	// Register custom validator for Product struct
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		v.RegisterTagNameFunc(func(fld reflect.StructField) string {
			name := strings.SplitN(fld.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
		v.RegisterStructValidation(api.ValidateProductRequest, types.Product{})
	}

	// Setup API server
	server := api.NewServer()
	server.SetupRouter()
	err = server.SetupServices(mongoClient)
	if err != nil {
		return nil, err
	}

	return server, nil
}
