package native_golang_endpoints

import (
	"fmt"
	"net/http"
)

func SimpleHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Custom-Header", "custom-value")
	fmt.Fprintf(w, "Hello World")
	printUrlThings(r)
}

func ComplexHelloWorld(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Custom-Header", "custom-value")
	fmt.Fprintf(w, "Complex Hello World")
	printUrlThings(r)
}

func printUrlThings(r *http.Request) {
	method := r.Method
	url := r.URL
	headers := r.Header
	queryParameters := r.URL.Query()

	fmt.Println(method, url.String(), headers.Values(""), queryParameters.Get("id"))
}
