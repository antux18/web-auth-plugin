package web_auth_plugin

import (
	"bytes"
	"context"
	"fmt"
	"net/http"
)

type Config struct {
	Headers map[string]string `json:"headers,omitempty"`
}

func CreateConfig() *Config {
	return &Config{
		Headers: make(map[string]string),
	}
}

type WebAuth struct {
	next     http.Handler
}

func New(ctx context.Context, next http.Handler, config *Config, name string) (http.Handler, error) {
	return &WebAuth{
		next:     next,
	}, nil
}

func (a *WebAuth) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	_, err := req.Cookie("authtoken")
	if err == http.ErrNoCookie {
		http.Redirect(rw, req, "/login", http.StatusFound)
		return
	}

	a.next.ServeHTTP(rw, req)
}
