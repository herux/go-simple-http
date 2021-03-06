package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/PuerkitoBio/goquery"
)

const AUTH_KEY = "1234567890"

func index(res http.ResponseWriter, req *http.Request) {
	file, err := os.Open("./templates/index.html")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()

	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}

	os.Setenv("AUTH", AUTH_KEY)
	s := os.Getenv("AUTH")
	doc.Find("input[name='auth']").SetAttr("value", s)
	sHtml, _ := doc.Html()
	res.Write([]byte(sHtml))
}

func upload(res http.ResponseWriter, req *http.Request) {
	req.ParseMultipartForm(10 << 20)

	file, handler, err := req.FormFile("data")
	if err != nil {
		fmt.Println("Error: ", err)
		return
	}
	defer file.Close()
	fmt.Println("file uploaded: ", handler.Filename)
	fmt.Println("size: ", handler.Size)
	fmt.Println("header: ", handler.Header.Values("Content-Type")[0])

	imageMimeType := "image/jpeg"
	if handler.Header.Values("Content-Type")[0] != imageMimeType {
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte("Invalid file type"))
		return
	}

	if req.FormValue("auth") != AUTH_KEY {
		res.WriteHeader(http.StatusForbidden)
		res.Write([]byte("Access denied"))
		return
	}

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
