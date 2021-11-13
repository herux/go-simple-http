package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func index(res http.ResponseWriter, req *http.Request) {
	http.ServeFile(res, req, "./templates/index.html")
}

func upload(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)

	file, handler, err := req.FormFile("file")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	fmt.Println("file uploaded: ", handler.Filename)
	fmt.Println("size: ", handler.Size)
	fmt.Println("header: ", handler.Header)

	tempFile, err := ioutil.TempFile("tempfiles", handler.Filename)
	if err != nil {
		fmt.Println(err)
	}
	defer tempFile.Close()

	// read all of the contents of our uploaded file into a
	// byte array
	fileBytes, err := ioutil.ReadAll(file)
	if err != nil {
		fmt.Println(err)
	}
	// write this byte array to our temporary file
	tempFile.Write(fileBytes)
	// return that we have successfully uploaded our file!
	fmt.Fprintf(res, "Successfully Uploaded File\n")
}

func main() {
	http.HandleFunc("/", index)
	http.HandleFunc("/upload", upload)

	port := ":8090"
	fmt.Println("Open http://localhost" + port)
	http.ListenAndServe(port, nil)
}
