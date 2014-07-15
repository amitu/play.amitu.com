package main

import (
	"net/http"
	"github.com/GeertJohan/go.rice"
)

func main() {
	fs := http.FileServer(rice.MustFindBox("public").HTTPBox())
	http.Handle("/", fs)
	http.ListenAndServe(":8877", &CORSMiddleware{http.DefaultServeMux})
}
