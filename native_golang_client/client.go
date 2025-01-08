package native_golang_client

import (
	"fmt"
	"io"
	"net/http"
)

func SimpleClient() {
	resp, err := http.Get("http://localhost:8080")

	if err != nil {
		fmt.Println("error")
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("Response:", string(body))
}
