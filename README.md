# Middleware: Request Body Parser

## How to use?
```
var api rest.API

api.Use(bodyparser.JSON())


// To get body
ctx.Body

// cast 
body := ctx.Body.(*orderedjson.OrderedMap)

```
