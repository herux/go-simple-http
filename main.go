package main

import (
	"fmt"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./templates/index.html")
}

func upload(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)

}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)

	port := ":8090"
	fmt.Println("Open http://localhost" + port)
	http.ListenAndServe(port, nil)
}
