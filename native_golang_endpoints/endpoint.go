package native_golang_endpoints

import (
	"fmt"
	"io"
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
	w.WriteHeader(http.StatusForbidden)
	fmt.Fprintf(w, "Complex Hello World")
	printUrlThings(r)
}

func SimpleHelloWithBody(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		fmt.Println("Error reading body")
	}

	fmt.Println("Body:", string(body))
}

func HandleRoot(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		SimpleHelloWorld(w, r)
	case http.MethodPost:
		SimpleHelloWithBody(w, r)
	default:
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
	}
}

func printUrlThings(r *http.Request) {
	method := r.Method
	url := r.URL
	headers := r.Header
	queryParameters := r.URL.Query()
	r.ParseForm()
	formData := r.Form

	fmt.Println(method, url.String(), headers.Values(""), queryParameters.Get("id"), formData)
}
