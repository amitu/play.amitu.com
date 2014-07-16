package main

import (
	"flag"
	"net/http"
	_ "code.google.com/p/go.tools/playground"
)

var (
	Agent string
)

type RepresentTransport struct {
	http.Transport
}

func (t *RepresentTransport) RoundTrip(req *http.Request) (*http.Response, error) { 
	req.Header.Set("User-Agent", Agent)
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
	flag.StringVar(&Agent, "agent", "represent", "User-Agent for play.golang.org")
	http.DefaultClient.Transport = &RepresentTransport{}
}
