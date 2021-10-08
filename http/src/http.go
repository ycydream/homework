package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"
)

func MyHandler(w http.ResponseWriter, r *http.Request) {
	header := r.Header
	for k, v := range header {
		//fmt.Println(k, v)
		w.Header().Set(k, strings.Join(v, ""))
	}
	version := os.Getenv("VERSION")
	fmt.Println(version)
	if version == "" {
		w.Header().Set("VERSION", "")
	} else {
		w.Header().Set("VERSION", version)
	}
	//fmt.Println("Header全部数据:", header)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":8080", nil)
}
