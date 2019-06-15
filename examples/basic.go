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
		ctx.JSON(ctx.Body)
	})

	fmt.Println("Starting server.")

	http.ListenAndServe(":8080", api)
}
