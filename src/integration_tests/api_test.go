package integration_tests

import (
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/corey888773/ztp-api/src/app"
	"github.com/corey888773/ztp-api/src/data"
	"github.com/corey888773/ztp-api/src/types"
	"github.com/stretchr/testify/assert"
)

func Test_ProductCreation(t *testing.T) {
	// Setup
	ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
	defer cancel()
	defer ClearDb(t)

	server, err := app.CreateApp(ctx)
	assert.NoError(t, err, "failed to create application")

	ts := httptest.NewServer(server.Router)
	defer ts.Close()

	testProduct := types.CreateProductRequest{
		Product: types.Product{
			Name:     "test",
			Category: types.Books,
			Quantity: 3,
			Price:    10,
		},
	}
	productJSON, err := json.Marshal(testProduct)

	// Test

	// CREATE PRODUCT
	resp, err := http.Post(ts.URL+"/api/v1/products", "application/json", strings.NewReader(string(productJSON)))
	assert.NoError(t, err, "failed to POST new product")
	defer resp.Body.Close()
	assert.Equal(t, http.StatusCreated, resp.StatusCode)

	// GET ALL PRODUCTS
	getAllResp, err := http.Get(ts.URL + "/api/v1/products")
	assert.NoError(t, err)
	assert.Equal(t, getAllResp.StatusCode, http.StatusOK)

	var allProducts []data.Product
	getAllBody, err := ioutil.ReadAll(getAllResp.Body)
	assert.NoError(t, err)
	err = json.Unmarshal(getAllBody, &allProducts)

	// GET PRODUCT BY ID
	getResp, err := http.Get(ts.URL + "/api/v1/products/" + allProducts[0].ID)
	assert.NoError(t, err, "failed to GET product by id")
	defer getResp.Body.Close()
	assert.Equal(t, http.StatusOK, getResp.StatusCode)

	getBody, err := ioutil.ReadAll(getResp.Body)
	assert.NoError(t, err, "failed to read GET response body")

	var fetchedProduct data.Product
	err = json.Unmarshal(getBody, &fetchedProduct)
	assert.NoError(t, err, "failed to unmarshal fetched product")
	assert.Equal(t, testProduct.Name, fetchedProduct.Name)
	assert.Equal(t, string(testProduct.Category), fetchedProduct.Category)
	assert.Equal(t, testProduct.Quantity, fetchedProduct.Quantity)
	assert.Equal(t, testProduct.Price, fetchedProduct.Price)
}
