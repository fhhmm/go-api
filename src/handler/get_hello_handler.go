package handler

import (
	"sample/src/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetHello(p operations.GetGreetingParams) middleware.Responder {
	return operations.NewGetGreetingOK().WithPayload("Hello Swagger")
}
