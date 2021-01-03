package main

import (
	"fmt"
	"net"
	"os"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "host")
		os.Exit(1)
	}
	service := os.Args[0]

	conn, err := net.Dial("ip4.icmp", service)
	checkErr(err)

	var msg [512]byte
	msg[0] = 8
	msg[1] = 0
	msg[2] = 0
	msg[3] = 0
	msg[4] = 0
	msg[5] = 13
	msg[6] = 0
	msg[7] = 37
	length := 8
	check := checkSum(msg[:length])
	msg[2] = byte(check >> 8)
	fmt.Println(conn)
}

func checkSum(msg []byte) uint16 {
	sum := 0
	n := 0
	for n+1 < len(msg) {
		sum += (int(msg[n]) << 8) | int(msg[n+1])
		n++
	}
	if n < len(msg) {
		sum += int(msg[n]) << 8
	}
	sum += (sum >> 16) + (sum & 0xffff)
	sum += sum >> 16
	return uint16(^sum)
}

func checkErr(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "failed error: %s", err.Error())
		os.Exit(1)
	}
}
