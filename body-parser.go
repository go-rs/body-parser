/*!
 * go-rs/body-parser
 * Copyright(c) 2019 Roshan Gade
 * MIT Licensed
 */
package bodyparser

import (
	"encoding/json"
	"errors"
	"net/url"
	"regexp"
	"strings"

	"github.com/go-rs/ordered-json"
	"github.com/go-rs/rest-api-framework"
)

/**
 * Basic level form data parsing
 */
func parseFormData(data url.Values) (body *orderedjson.OrderedMap) {
	//TODO: extended parsing
	if len(data) > 0 {
		body := orderedjson.NewOrderedMap()
		for key, val := range data {
			if len(val) > 1 {
				body.Set(key, val)
			} else {
				body.Set(key, data.Get(key))
			}
		}
	}
	return
}

/**
 * Body Parser
 */
func JSON() rest.Handler {

	formHeader := regexp.MustCompile(`^multipart/form-data`)

	return func(ctx *rest.Context) {
		if ctx.Request.Method == "GET" {
			return
		}

		contentType := strings.ToLower(ctx.Request.Header.Get("content-type"))

		var body = orderedjson.NewOrderedMap()

		if contentType == "application/json" {
			err := json.NewDecoder(ctx.Request.Body).Decode(&body)
			if err != nil {
				ctx.Status(400).Throw(errors.New("MALFORMED_BODY"))
				return
			}
		} else if contentType == "application/x-www-form-urlencoded" {
			err := ctx.Request.ParseForm()
			if err != nil {
				ctx.Status(400).Throw(errors.New("FORM_PARSE_ERROR"))
				return
			}

			body = parseFormData(ctx.Request.PostForm)
		} else if formHeader.MatchString(contentType) {
			err := ctx.Request.ParseMultipartForm(2000)
			if err != nil {
				ctx.Status(400).Throw(errors.New("MULTIPART_FORM_PARSE_ERROR"))
				return
			}

			body = parseFormData(ctx.Request.PostForm)
		}
		ctx.Body = body
	}
}
