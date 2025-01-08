package native_golang_endpoints

import (
	"fmt"
	"net/http"
)

func SimpleHelloWorld() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Hello World")
	})
}
