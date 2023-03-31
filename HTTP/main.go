package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// resp is a struct, Body is of type io.ReadCloser
	// ReadCloser is an interface with a Reader and Closer
	// Reader and closer are interfaces
	resp, err := http.Get("http://google.com")
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	fmt.Println((resp))
}