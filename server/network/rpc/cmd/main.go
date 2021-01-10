package main

import (
	"log"
	"net"
	"net/http"
	"net/rpc"
	"os"
	"os/signal"
	"server/server/network/rpc/server"
	"syscall"
)

func main() {
	greet := new(server.Greeter)
	err := rpc.Register(greet)
	if err != nil {
		log.Fatal(err)
	}
	rpc.HandleHTTP()
	listen, err := net.Listen("tcp", ":8081")
	if err != nil {
		log.Fatal(err)
	}
	go http.Serve(listen, nil)
	quite := make(chan os.Signal, 1)
	signal.Notify(quite, syscall.SIGINT, syscall.SIGTERM)
	<-quite
}
