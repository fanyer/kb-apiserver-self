package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/cepave", func(w http.ResponseWriter, req *http.Request) {
		req.ParseForm()
		if req.Method == "GET" || req.Method == "POST" {
			fmt.Println(req.ContentLength)
			firstname := req.FormValue("firstname")
			lastname := req.FormValue("lastname")
			w.Write([]byte(fmt.Sprintf("[%s] Hello, %s %s!", req.Method, firstname, lastname)))
		} else {
			http.Error(w, "The method is not allowed.", http.StatusMethodNotAllowed)
		}
	})

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("ListenAndServe failed: ", err)
	}
}
