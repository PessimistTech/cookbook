package handlers

import (
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

type APIError struct {
	StatusCode int         `json:"status"`
	Message    string      `json:"msg"`
	Err        interface{} `json:"err"`
}

func NewAPIError(statusCode int, message string, err interface{}) *APIError {
	return &APIError{
		StatusCode: statusCode,
		Message:    message,
		Err:        err,
	}
}

func (e *APIError) Error() string {
	return e.Message
}

func ErrorHandler() gin.HandlerFunc {
	return ErrorHandlerLogic(gin.ErrorTypeAny)
}

func ErrorHandlerLogic(errType gin.ErrorType) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		errList := ctx.Errors.ByType(errType)

		if len(errList) > 0 {
			var parsedErr *APIError
			err := errList[0].Err
			switch err.(type) {
			case *APIError:
				parsedErr = err.(*APIError)
			default:
				parsedErr = &APIError{
					StatusCode: 500,
					Message:    "internal server error",
					Err:        "unable to process request",
				}
			}
			logrus.WithError(parsedErr).Errorf("failed request to %s", ctx.FullPath())
			ctx.JSON(parsedErr.StatusCode, parsedErr)
			ctx.Abort()
		}
	}
}
