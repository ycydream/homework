package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"strings"
)

func HealthzHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(200)
	fmt.Fprintf(w, "200")
}

func CommonHandler(w http.ResponseWriter, r *http.Request) {
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
	ip := ClientPublicIP(r)
	if ip == "" {
		ip = ClientIP(r)
	}
	w.WriteHeader(200)
	fmt.Printf("IP地址:%s；HTTP 返回码：%s\n", ip, "200")
	//fmt.Println("Header全部数据:", header)
	fmt.Fprintf(w, "Hello World")
}

func main() {
	http.HandleFunc("/healthz", HealthzHandler)
	http.HandleFunc("/", CommonHandler)
	http.ListenAndServe(":80", nil)
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}

func ClientPublicIP(r *http.Request) string {
	var ip string
	for _, ip = range strings.Split(r.Header.Get("X-Forwarded-For"), ",") {
		ip = strings.TrimSpace(ip)
		if ip != "" {
			return ip
		}
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
