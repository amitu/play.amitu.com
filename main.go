package main

import (
	"net/http"
)

func main() {
	http.ListenAndServe(":8877", &CORSMiddleware{http.DefaultServeMux})
}
