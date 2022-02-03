package swaggerui_test

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"gotest.tools/v3/assert"

	swaggerui "go.skymeyer.dev/swagger-ui-bindata"
)

var (
	mainJSCallRegion                = "// Begin Swagger UI call region"
	mainDefaultSpec                 = "https://petstore.swagger.io/v2/swagger.json"
	mainEmbeddedSpec                = "./spec.json"
	distSwaggerUICSS                = "sourceMappingURL=swagger-ui.css.map"
	distSwaggerUIBundleVersion      = `Vr="4.4.1"`
	distSwaggerUIBundleJS           = "For license information please see swagger-ui-bundle.js.LICENSE.txt"
	distSwaggerUIStandalonePresetJS = "For license information please see swagger-ui-standalone-preset.js.LICENSE.txt"
)

func TestExpectedFiles(t *testing.T) {
	for _, d := range []struct {
		test     string // test description
		prefix   string // site/mux prefix
		url      string // URL to call
		status   int    // expected HTTP status code
		contains string // expected returned content
		opts     []swaggerui.Option
	}{
		{
			test:     "Test version",
			prefix:   "/",
			url:      "/dist/swagger-ui-bundle.js",
			status:   http.StatusOK,
			contains: distSwaggerUIBundleVersion,
			opts:     []swaggerui.Option{},
		},
		{
			test:     "Test main index page with defaults",
			prefix:   "/",
			url:      "/",
			status:   http.StatusOK,
			contains: mainJSCallRegion,
			opts:     []swaggerui.Option{},
		},
		{
			test:     "Test main index page default spec",
			prefix:   "/",
			url:      "/",
			status:   http.StatusOK,
			contains: mainDefaultSpec,
			opts:     []swaggerui.Option{},
		},
		{
			test:     "Test /dist/swagger-ui.css with defaults",
			prefix:   "/",
			url:      "/dist/swagger-ui.css",
			status:   http.StatusOK,
			contains: distSwaggerUICSS,
			opts:     []swaggerui.Option{},
		},
		{
			test:     "Test /dist/swagger-ui-bundle.js with defaults",
			prefix:   "/",
			url:      "/dist/swagger-ui-bundle.js",
			status:   http.StatusOK,
			contains: distSwaggerUIBundleJS,
			opts:     []swaggerui.Option{},
		},
		{
			test:     "Test /dist/swagger-ui-standalone-preset.js with defaults",
			prefix:   "/",
			url:      "/dist/swagger-ui-standalone-preset.js",
			status:   http.StatusOK,
			contains: distSwaggerUIStandalonePresetJS,
			opts:     []swaggerui.Option{},
		},
		{
			test:     "Test main index page with prefix",
			prefix:   "/foo/",
			url:      "/foo/",
			status:   http.StatusOK,
			contains: mainJSCallRegion,
			opts: []swaggerui.Option{
				swaggerui.WithPrefix("/foo/"),
			},
		},
		{
			test:     "Test main index page with prefix (redirect)",
			prefix:   "/foo/bar/",
			url:      "/foo/bar",
			status:   http.StatusMovedPermanently,
			contains: "",
			opts: []swaggerui.Option{
				swaggerui.WithPrefix("/foo/bar/"),
			},
		},
		{
			test:     "Test /dist/swagger-ui.css with prefix",
			prefix:   "/foo/bar/",
			url:      "/foo/bar/dist/swagger-ui.css",
			status:   http.StatusOK,
			contains: distSwaggerUICSS,
			opts: []swaggerui.Option{
				swaggerui.WithPrefix("/foo/bar/"),
			},
		},
		{
			test:     "Test /dist/swagger-ui-bundle.js with prefix",
			prefix:   "/foo/bar/",
			url:      "/foo/bar/dist/swagger-ui-bundle.js",
			status:   http.StatusOK,
			contains: distSwaggerUIBundleJS,
			opts: []swaggerui.Option{
				swaggerui.WithPrefix("/foo/bar/"),
			},
		},
		{
			test:     "Test /dist/swagger-ui-standalone-preset.js with prefix",
			prefix:   "/foo/bar/",
			url:      "/foo/bar/dist/swagger-ui-standalone-preset.js",
			status:   http.StatusOK,
			contains: distSwaggerUIStandalonePresetJS,
			opts: []swaggerui.Option{
				swaggerui.WithPrefix("/foo/bar/"),
			},
		},
		{
			test:     "Test main index page with spec URL",
			prefix:   "/",
			url:      "/",
			status:   http.StatusOK,
			contains: "http://foo.bar.com/openapi.json",
			opts: []swaggerui.Option{
				swaggerui.WithSpecURL("http://foo.bar.com/openapi.json"),
			},
		},
		{
			test:     "Test main index page with embedded spec",
			prefix:   "/",
			url:      "/",
			status:   http.StatusOK,
			contains: mainEmbeddedSpec,
			opts: []swaggerui.Option{
				swaggerui.WithEmbeddedSpec([]byte("fakespec")),
			},
		},
		{
			test:     "Test main index page with embedded spec (embedded spec content)",
			prefix:   "/",
			url:      "/spec.json",
			status:   http.StatusOK,
			contains: "fakespec",
			opts: []swaggerui.Option{
				swaggerui.WithEmbeddedSpec([]byte("fakespec")),
			},
		},
		{
			test:     "Test main index page with embedded spec and prefix (embedded spec content)",
			prefix:   "/foo/bar/",
			url:      "/foo/bar/spec.json",
			status:   http.StatusOK,
			contains: "fakespec",
			opts: []swaggerui.Option{
				swaggerui.WithPrefix("/foo/bar/"),
				swaggerui.WithEmbeddedSpec([]byte("fakespec")),
			},
		},
	} {
		r := httptest.NewRequest(http.MethodGet, d.url, nil)
		w := httptest.NewRecorder()

		mux := http.NewServeMux()
		mux.Handle(d.prefix, swaggerui.New(d.opts...).Handler())
		mux.ServeHTTP(w, r)

		assert.Equal(t, d.status, w.Code, d.test)
		assert.Assert(t, strings.Contains(w.Body.String(), d.contains), d.test)
	}
}
