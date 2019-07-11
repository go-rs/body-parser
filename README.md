# Middleware: Request Body Parser

## How to use?
```
var api = rest.New("/")

api.Use(bodyparser.JSON(memorySize int64))


// To get body
ctx.Body

```
