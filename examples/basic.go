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
	api.Use(bodyparser.Load())

	api.All("/", func(ctx *rest.Context) {
		body, _ := ctx.Get("body")
		ctx.JSON(body)
	})

	fmt.Println("Starting server.")

	http.ListenAndServe(":8080", api)
}
