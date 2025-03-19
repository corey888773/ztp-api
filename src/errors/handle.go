package custom_errors

import (
	"errors"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/mongo"
)

func Handle(ctx *gin.Context, err error) {
	log.Println(err)

	if errors.Is(err, mongo.ErrNoDocuments) {
		WithError(ctx, RecordNotFound, http.StatusNotFound)
		return
	}

	var writeException mongo.WriteException
	if errors.As(err, &writeException) {
		for _, we := range writeException.WriteErrors {
			if mongo.IsDuplicateKeyError(we) {
				WithError(ctx, DuplicateRecord, http.StatusBadRequest)
				return
			}
		}
	}

	var validationErrs validator.ValidationErrors
	if errors.As(err, &validationErrs) {
		var msg []string
		for _, fieldErr := range validationErrs {
			msg = append(msg, fieldErr.Error())
		}
		WithErrorMessage(ctx, strings.Join(msg, "; "), http.StatusBadRequest)
		return
	}

	if strings.Contains(err.Error(), "binding") {
		WithError(ctx, InvalidInput, http.StatusBadRequest)
		return
	}

	WithError(ctx, InternalServerError, http.StatusInternalServerError)
}

type CustomErrorMessages string

const (
	RecordNotFound      CustomErrorMessages = "Record not found"
	DuplicateRecord     CustomErrorMessages = "Duplicate record"
	InvalidInput        CustomErrorMessages = "Invalid input"
	InternalServerError CustomErrorMessages = "Internal server error"
)

func WithError(ctx *gin.Context, message CustomErrorMessages, status int) {
	ctx.JSON(status, gin.H{"error": message})
	return
}

func WithErrorMessage(ctx *gin.Context, message string, status int) {
	ctx.JSON(status, gin.H{"error": message})
	return
}
