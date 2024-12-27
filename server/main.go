package main

import "net/http"

const address = "localhost:8080"

func main() {
	_ = http.NewServeMux()

}
