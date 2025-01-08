package native_golang_endpoints

import (
	"fmt"
	"net/http"
)

func SimpleHelloWorld() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	}
}

func ComplexHelloWorld() func(http.ResponseWriter, *http.Request) {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Complex Hello World")
	}
}
