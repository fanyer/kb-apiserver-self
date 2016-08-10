package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

func main() {
	// POST
	resp, err := http.PostForm("http://localhost:9000/cepave",
		url.Values{"firstname": {"Kordan"}, "lastname": {"Ou"}})
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("POST OK: ", string(body), resp)
	}

	// GET
	resp, err = http.Get("http://localhost:9000/cepave?firstname=Kordan&lastname=Ou")
	if err != nil {
		fmt.Println(err)
	} else {
		body, _ := ioutil.ReadAll(resp.Body)
		fmt.Println("GET OK: ", string(body), resp)
	}
}
