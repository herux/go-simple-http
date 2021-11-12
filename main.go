package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./templates/index.html")
}

func main() {
	http.HandleFunc("/", index)

	port := ":8090"
	fmt.Println("Open http://localhost" + port)
	http.ListenAndServe(port, nil)
}
