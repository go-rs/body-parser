package main

import (
	"fmt"
	"net/http"

	bodyparser "github.com/go-rs/body-parser"
	orderedjson "github.com/go-rs/ordered-json"

	"github.com/go-rs/rest-api-framework"
)

func main() {
	var api rest.API

	// request interceptor / middleware
	api.Use(bodyparser.JSON())

	api.All("/", func(ctx *rest.Context) {
		//prettyBytes, _ := json.Marshal(ctx.Body)
		//ctx.SetHeader("content-type", "application/json")
		//ctx.Write(prettyBytes)

		body := ctx.Body.(*orderedjson.OrderedMap)

		ctx.JSON(body)
	})

	fmt.Println("Starting server.")

	http.ListenAndServe(":8080", &api)
}
