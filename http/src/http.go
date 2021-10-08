package main

import (
	"fmt"
	"net/http"
	"strings"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k, v := range header {
		//fmt.Println(k, v)
		w.Header().Set(k, strings.Join(v, ""))
	}
	//fmt.Println("Header全部数据:", header)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":8080", nil)
}
