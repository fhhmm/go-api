package handler

import (
	"sample/src/gen/restapi/operations"

	"github.com/go-openapi/runtime/middleware"
)

func GetHello(p operations.GetGreetingParams) middleware.Responder {
	// p.Nameは*string型なので、文字列を取得したいときは*をつけてデリファレンスする
	return operations.NewGetGreetingOK().WithPayload(*p.Name)
}

func PostHello(p operations.PostGreetingParams) middleware.Responder {
	return operations.NewGetGreetingOK().WithPayload(*p.Greeting.Name)
}
