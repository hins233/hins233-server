package main

import (
	"fmt"
	"log"
	"net/rpc"
	"server/server/network/rpc/server"
)

func main() {
	testRpc()
}

func testRpc() {
	client, err := rpc.DialHTTP("tcp", "127.0.0.1:8081")
	if err != nil {
		log.Fatal("dialing failed", err)
	}
	// rpc 调用不能传 nil
	greet := new(server.Greet)
	var answer server.Answer
	*greet = "hins"
	err = client.Call("Greeter.SayHello", greet, &answer)
	if err != nil {
		log.Fatal("SayHello rpc failed", err)
	}
	fmt.Println(answer)
	args := &server.Args{9, 9}
	var reply int
	multiCall := client.Go("Greeter.Multiply", args, &reply, nil)
	fmt.Println(reply)
	call := <-multiCall.Done
	fmt.Println(call.Reply)
	fmt.Println(reply)
}
