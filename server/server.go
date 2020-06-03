package main

import (
	"fmt"
	"log"
	"net/http"
)

var types map[string]bool

func main() {
	types = make(map[string]bool)
	types[".py"] = true
	types[".go"] = true
	types[".c"] = true
	types[".cpp"] = true
	types[".js"] = true
	types[".ts"] = true
	types[".rs"] = true
	types[".java"] = true
	port := ":5000"
	http.Handle("/", http.FileServer(http.Dir("../build/")))
	http.HandleFunc("/api", index)
	fmt.Println("Listening on " + port)
	log.Fatal(http.ListenAndServe(port, nil))
}
