// go-rs/body-parser
// Copyright(c) 2019 Roshan Gade. All rights reserved.
// MIT Licensed

package bodyparser

import (
	"encoding/json"
	"net/http"
	"net/url"
	"strings"

	"github.com/go-rs/rest-api-framework"
)

// parse form data into map
func parseFormData(data url.Values) (body map[string]interface{}) {
	//TODO: extended parsing
	if len(data) > 0 {
		body = make(map[string]interface{})
		for key, val := range data {
			if len(val) > 1 {
				body["key"] = val
			} else {
				body["key"] = data.Get(key)
			}
		}
	}
	return
}

// Error codes on request body parse
const (
	ErrCodeRequestSize    = "REQUEST_SIZE_EXCEED"
	ErrCodeMalformedBody  = "MALFORMED_BODY"
	ErrCodeFormParse      = "FORM_PARSE_ERROR"
	ErrCodeMultiformParse = "MULTIPART_FORM_PARSE_ERROR"
)

// Read requested JSON body and store it into context in map[string]interface{} format
func JSON(maxMemory int64) rest.Handler {

	return func(ctx *rest.Context) {
		if ctx.Request.Method == http.MethodGet || ctx.Request.Method == http.MethodOptions {
			return
		}

		if maxMemory < ctx.Request.ContentLength {
			ctx.Status(http.StatusPreconditionFailed).Throw(ErrCodeRequestSize)
			return
		}

		contentType := strings.ToLower(ctx.Request.Header.Get("content-type"))

		var body map[string]interface{}

		if strings.Contains(contentType, "application/json") {
			err := json.NewDecoder(ctx.Request.Body).Decode(&body)
			if err != nil {
				ctx.Status(http.StatusBadRequest).ThrowWithError(ErrCodeMalformedBody, err)
				return
			}
		} else if strings.Contains(contentType, "application/x-www-form-urlencoded") {
			err := ctx.Request.ParseForm()
			if err != nil {
				ctx.Status(http.StatusBadRequest).ThrowWithError(ErrCodeFormParse, err)
				return
			}

			body = parseFormData(ctx.Request.PostForm)
		} else if strings.Contains(contentType, "multipart/form-data") {
			err := ctx.Request.ParseMultipartForm(maxMemory)
			if err != nil {
				ctx.Status(http.StatusBadRequest).ThrowWithError(ErrCodeMultiformParse, err)
				return
			}

			body = parseFormData(ctx.Request.PostForm)
		}
		ctx.Body = body
	}
}
