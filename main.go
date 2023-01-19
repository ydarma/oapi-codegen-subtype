package main

import (
	"context"
	"oapi-codegen-subtype/api"
	"oapi-codegen-subtype/models"
)

type handler struct {
}

func (h handler) Test(ctx context.Context, request api.TestRequestObject) (api.TestResponseObject, error) {
	return api.Test200JSONResponse(models.Foo{}), nil
}

func main() {
	_, err := api.GetSwagger()
	if err != nil {
		panic(err)
	}
}
