package api

import (
	"github.com/corey888773/ztp-api/src/types"
	"github.com/go-playground/validator/v10"
)

/* Validator Rules */

func ValidateProductRequest(sl validator.StructLevel) {
	product, ok := sl.Current().Interface().(types.Product)
	if !ok {
		return
	}

	switch product.Category {
	case types.Electronics:
		if ok := validatePriceForElectronics(product.Price); !ok {
			sl.ReportError(product.Price, "Price", "price", "gte", "100")
		}
	case types.Books:
		if ok := validatePriceForBooks(product.Price); !ok {
			sl.ReportError(product.Price, "Price", "price", "range", "5-500")
		}
	case types.Clothing:
		if ok := validatePriceForClothing(product.Price); !ok {
			sl.ReportError(product.Price, "Price", "price", "range", "10-5000")
		}
	default:
	}
}

/* Validation Functions */

func validatePriceForElectronics(price float64) bool {
	if price < 100 {
		return false
	}
	return true
}

func validatePriceForBooks(price float64) bool {
	if price < 5 || price > 500 {
		return false
	}

	return true
}

func validatePriceForClothing(price float64) bool {
	if price < 10 || price > 5000 {
		return false
	}
	return true
}

func ShouldNotBeEmpty(s string) (bool, string) {
	if s == "" {
		return false, "Parameter should not be empty"
	}
	return true, ""
}
