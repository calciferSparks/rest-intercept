package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

const (
	exitSuccess    = iota
	exitServeError = iota
)

func main() {
	var exitCode int = program()
	if exitCode != 0 {
		os.Exit(exitCode)
	}
}

func program() int {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":80", nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, "hosting error:", err.Error())
	}
	return exitSuccess
}

func handler(writer http.ResponseWriter, request *http.Request) {
	fmt.Println("Received request:", request.URL.String())
	fmt.Println("Compression:", request.Header["Accept-Encoding"])
	io.Copy(os.Stdout, request.Body)
}
