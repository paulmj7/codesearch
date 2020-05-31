package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func URLParams(url string) (string, string) {
	var name, account string
	url, name = popParam(url)
	url, account = popParam(url)
	return name, account
}

func popParam(url string) (string, string) {
	i := strings.LastIndex(url, "/")
	return url[0:i], url[i+1:]
}

func SHA(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	data, er := ioutil.ReadAll(res.Body)
	if er != nil {
		log.Fatal(er)
	}

	var raw map[string]interface{}
	json.Unmarshal(data, &raw)
	return raw["commit"].(map[string]interface{})["sha"].(string)
}

func tree(url string) map[string]bool {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer res.Body.Close()

	data, er := ioutil.ReadAll(res.Body)
	if er != nil {
		log.Fatal(er)
	}

	var raw map[string]interface{}
	json.Unmarshal(data, &raw)
	test, _ := json.Marshal(raw["tree"])
	var tree []map[string]interface{}
	json.Unmarshal(test, &tree)

	paths := make(map[string]bool)

	for _, r := range tree {
		if r["type"] == "blob" && allowedFType(r["path"].(string)) {
			paths[r["path"].(string)] = true
		}
	}

	return paths
}

func allowedFType(path string) bool {
	i := strings.LastIndex(path, ".")
	if i == -1 {
		return false
	}
	ext := path[i:]
	if _, ok := types[ext]; ok {
		return true
	}

	return false
}

func file(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	data, er := ioutil.ReadAll(res.Body)
	if er != nil {
		log.Fatal(er)
	}

	return string(data)
}

func contents(url string) map[string]string {
	c := make(map[string]string)

	name, account := URLParams(url)
	sha := SHA("http://api.github.com/repos/" + account + "/" + name + "/branches/master")
	tree := tree("http://api.github.com/repos/" + account + "/" + name + "/git/trees/" + sha + "?recursive=1")

	for t := range tree {
		c[t] = file("http://raw.githubusercontent.com/" + account + "/" + name + "/master/" + t)
	}

	return c
}
