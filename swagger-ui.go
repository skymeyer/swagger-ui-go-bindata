//go:generate go-bindata -fs -prefix swagger-ui/dist/ -o bindata/bindata.go -pkg bindata swagger-ui/dist/...
package swaggerui

import (
	"bytes"
	"net/http"
	"path"
	"sync"

	"go.skymeyer.dev/swagger-ui-bindata/bindata"
)

var (
	indexPage    = "index.html"
	staticDir    = "dist/"
	embeddedSpec = "spec.json"
	replacePath  = "./"
	replaceSpec  = "https://petstore.swagger.io/v2/swagger.json"
)

// New createa a new SwaggerUI instance.
func New(opts ...Option) *SwaggerUI {
	s := &SwaggerUI{
		prefix: "/",
	}
	for _, opt := range opts {
		opt(s)
	}
	return s
}

// SwaggerUI represents a bindata based swagger-ui website handler.
type SwaggerUI struct {
	init      sync.Once
	index     []byte
	specEmbed []byte
	specURL   string
	prefix    string
}

// Handler returns the swagger-ui http.Handler
func (s *SwaggerUI) Handler() http.Handler {
	mux := http.NewServeMux()

	// Use swagger-ui-dist static files
	bindataDir := path.Join(s.prefix, staticDir) + "/"
	mux.Handle(bindataDir, http.StripPrefix(bindataDir, http.FileServer(bindata.AssetFile())))

	// Embedded spec support
	if s.specEmbed != nil {
		mux.HandleFunc(path.Join(s.prefix, embeddedSpec), func(w http.ResponseWriter, r *http.Request) {
			w.Write(s.specEmbed)
		})
	}

	// Index handler
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Customize original index.html
		s.init.Do(func() {
			index := bindata.MustAsset(indexPage)

			// Fixup resource paths
			index = bytes.ReplaceAll(index, []byte(replacePath), []byte("./"+staticDir))

			// Use embedded spec if requested
			if s.specEmbed != nil {
				index = bytes.ReplaceAll(index, []byte(replaceSpec), []byte("./"+embeddedSpec))
			}

			// Fallback to use spec by URL
			if s.specURL != "" {
				index = bytes.ReplaceAll(index, []byte(replaceSpec), []byte(s.specURL))
			}

			s.index = index
		})
		w.Write(s.index)
	})
	return mux
}
