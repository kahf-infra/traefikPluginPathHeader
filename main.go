package pathHeader

import (
	"context"
	"net/http"
)

type Config struct{}

func CreateConfig() *Config {
	return &Config{}
}

type PathHeader struct {
	next http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &PathHeader{next: next}, nil
}

func (p *PathHeader) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	req.Header.Set("X-Forwarded-Path", req.URL.Path)
	p.next.ServeHTTP(rw, req)
}
