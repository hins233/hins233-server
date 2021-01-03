package main

import (
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
)

var addr = "127.0.0.1:2002"

func main() {
	rs1 := "http://127.0.0.1:2003/base"
	url1, err := url.Parse(rs1)
	if err != nil {
		log.Panicln(err)
	}
	// 获取一个实例的代理
	proxy := httputil.NewSingleHostReverseProxy(url1)
	log.Println("Starting http server at " + addr)
	log.Fatal(http.ListenAndServe(addr, proxy))
}
