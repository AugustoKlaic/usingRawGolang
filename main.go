package main

import (
	"net/http"
	. "usingRawGolang/native_golang_endpoints"
)

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", HandleRoot)
	mux.HandleFunc("/complex/", ComplexHelloWorld)

	http.ListenAndServe(":8080", mux)
}
