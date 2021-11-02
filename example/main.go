package main

import (
	"net/http"

	swaggerui "go.skymeyer.dev/swagger-ui-bindata"
)

func main() {
	http.ListenAndServe(":8080", swaggerui.New(
		swaggerui.WithSpecURL("https://raw.githubusercontent.com/APIs-guru/openapi-directory/main/APIs/circleci.com/v1/openapi.yaml"),
	).Handler())
}
