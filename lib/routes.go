package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	setupResponse(&w, r)
	if (*r).Method == "OPTIONS" {
		return
	}

	if r.Method != "POST" {
		http.Error(w, "405 Method Not Allowed", http.StatusMethodNotAllowed)
	}

	decoder := json.NewDecoder(r.Body)
	var req ReqBody
	err := decoder.Decode(&req)
	if err != nil {
		fmt.Println(err)
		return
	}

	c := contents(req.URL)
	js, er := json.Marshal(c)
	if er != nil {
		fmt.Println(er)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
}

func setupResponse(w *http.ResponseWriter, r *http.Request) {
	(*w).Header().Set("Access-Control-Allow-Origin", "*")
	(*w).Header().Set("Access-Control-Allow-ZMethods", "POST, GET, OPTIONS, PUT, DELETE")
	(*w).Header().Set("Access-Control-Allow-Headers", "Accept, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
}
