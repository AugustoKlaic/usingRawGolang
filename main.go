package main

import (
	"net/http"
	. "usingRawGolang/native_golang_endpoints"
)

func main() {
	SimpleHelloWorld()
	http.ListenAndServe(":8080", nil)
}
