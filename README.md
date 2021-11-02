# Swagger UI Go Bindata

[![go-doc](https://pkg.go.dev/badge/go.skymeyer.dev/swagger-ui-bindata?status.svg)](https://pkg.go.dev/go.skymeyer.dev/swagger-ui-bindata)
[![Go Report Card](https://goreportcard.com/badge/go.skymeyer.dev/swagger-ui-bindata)](https://goreportcard.com/report/go.skymeyer.dev/swagger-ui-bindata)
[![GitHub license](https://img.shields.io/github/license/skymeyer/swagger-ui-go-bindata)](https://github.com/skymeyer/swagger-ui-go-bindata/blob/main/LICENSE)

This package embeds [swagger-ui-dist](https://github.com/swagger-api/swagger-ui/tree/master/dist) using
[go-bindata](https://github.com/go-bindata/go-bindata) and exposes an `http.Handler` to integrate in existing HTTP server projects. The generated `bindata` is available at `go.skymeyer.dev/swagger-ui-bindata/bindata` and is generated with the `-fs` option exposing an `http.FileSystem` interface.

## Usage

For the default Swagger UI, follow below code sample:

```go
package main

import (
	"net/http"
	swaggerui "go.skymeyer.dev/swagger-ui-bindata"
)

func main() {
	http.ListenAndServe(":8080", swaggerui.New().Handler())
}
```

More useful is the ability to set the OpenAPI spec you want to display instead of the default demo Petstore. This can be achieved by using the `swaggerui.WithSpecURL` option.

```go
swaggerui.New().Handler(
    swaggerui.WithSpecURL("https://foobar.com/openapi.yaml"),
)
```

Alternatively, the OpenAPI spec can also be passed in `[]byte` format if you wish to embed it directly in the Go binary. For this, use the `swaggerui.WithEmbeddedSpec` option.

```go
swaggerui.New().Handler(
    swaggerui.WithEmbeddedSpec(bindata.MustAsset("openapi.yaml")),
)
```

If both options are specified, the `WithEmbeddedSpec` has precedence.

## Example

Run the [example code](example/main.go) using `go run ./example` and point your browser to [http://localhost:8080/](http://localhost:8080/).
