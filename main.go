package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strings"
)

func main() {

	server := http.NewServeMux()
	server.HandleFunc("GET /", HandleRequest)

	if err := http.ListenAndServe(":8080", server); err != nil {
		fmt.Println("Server error:", err)
	}

}

func HandleRequest(res http.ResponseWriter, req *http.Request) {
	path := req.URL.Path
	// log.Println(path + " ")
	if path == "/" {
		path = "/index.html"
	}
	working_directory, err := os.Getwd()
	if err != nil {
		log.Fatalln("Error getting current working Directory")
		return
	}
	file_path := filepath.Join(working_directory, "www", path)
	file, err := os.OpenFile(file_path, os.O_RDONLY, os.FileMode('r'))
	if err != nil {
		res.WriteHeader(http.StatusNotFound)
		fmt.Fprintf(res, "404! Not Found!")
		return
	}
	content, err := io.ReadAll(file)
	if err != nil {
		log.Fatalln("no content")
		return
	}
	var content_type string
	extension := strings.Split(file.Name(), ".")[1]
	switch extension {
	case "js":
		content_type = "text/js"
	case "css":
		content_type = "text/css"

	case "html":
		content_type = "text/html"
	case "json":
		content_type = "application/json"
	default:
		content_type = "text/plain"
	}
	// http.DetectContentType(content)
	res.Header().Set("Content-Type", content_type)
	res.WriteHeader(http.StatusOK)
	res.Write(content)
}
