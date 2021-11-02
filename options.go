package swaggerui

// Option is a SwaggerUI configuration option.
type Option func(*SwaggerUI)

// WithSpecURL sets the OpenAPI spec URL.
func WithSpecURL(url string) Option {
	return func(s *SwaggerUI) {
		s.specURL = url
	}
}

// WithEmbeddedSpec sets the OpenAPI spec byte content.
func WithEmbeddedSpec(spec []byte) Option {
	return func(s *SwaggerUI) {
		s.specEmbed = spec
	}
}

// WithPrefix sets the http mux prefix.
func WithPrefix(prefix string) Option {
	return func(s *SwaggerUI) {
		s.prefix = prefix
	}
}
