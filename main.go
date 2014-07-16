package main

import (
	"log"
	"flag"
	"net/http"
	"github.com/GeertJohan/go.rice"
)

var (
	hostport string
)

func main() {
	flag.StringVar(&hostport, "http", ":8877", "HTTP host:port")
	flag.Parse()

	fs := http.FileServer(rice.MustFindBox("public").HTTPBox())
	http.Handle("/", fs)

	log.Printf("Started webserver at %s.", hostport)
	http.ListenAndServe(hostport, &CORSMiddleware{http.DefaultServeMux})
}
