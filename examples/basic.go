package main

import (
	"fmt"
	"net/http"

	bodyparser "github.com/go-rs/body-parser"

	"github.com/go-rs/rest-api-framework"
)

func main() {
	var api rest.API

	// request interceptor / middleware
	api.Use(bodyparser.JSON(2000))

	api.All("/", func(ctx *rest.Context) {
		body := ctx.Body
		ctx.JSON(body)
	})

	api.OnErrors([]string{bodyparser.ErrCodeMalformedBody, bodyparser.ErrCodeFormParse, bodyparser.ErrCodeMultiformParse}, func(ctx *rest.Context) {
		ctx.Status(400).JSON(`{"message": "Malformed body"}`)
	})

	fmt.Println("Starting server.")

	http.ListenAndServe(":8080", &api)
}
