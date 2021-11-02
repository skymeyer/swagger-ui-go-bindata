package main

import (
	"net/http"

	swaggerui "go.skymeyer.dev/swagger-ui-bindata"
)

func main() {
	mux := http.NewServeMux()
	mux.Handle("/swagger-ui/", swaggerui.New(
		swaggerui.WithPrefix("/swagger-ui/"),
		swaggerui.WithSpecURL("https://raw.githubusercontent.com/APIs-guru/openapi-directory/main/APIs/circleci.com/v1/openapi.yaml"),
	).Handler())
	http.ListenAndServe(":8080", mux)
}
