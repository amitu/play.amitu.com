package main

import (
	"net/http"
	_ "code.google.com/p/go.tools/playground"
)

type RepresentTransport struct {
	http.Transport
}

func (t *RepresentTransport) RoundTrip(req *http.Request) (*http.Response, error) { 
	req.Header.Set("User-Agent", "represent")
	return t.Transport.RoundTrip(req)
}

type CORSMiddleware struct {
	handler http.Handler
}

func (m *CORSMiddleware) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/compile" || r.URL.Path == "/share" {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set(
			"Access-Control-Allow-Methods", "GET, PUT, POST, DELETE, OPTIONS",
		)
		w.Header().Set(
			"Access-Control-Allow-Headers",
			"Content-Type, Authorization, X-Requested-With",
		)
		w.Header().Set("Content-Type", "application/json")
	}
	m.handler.ServeHTTP(w, r)
}

func init() {
	http.DefaultClient.Transport = &RepresentTransport{}
}