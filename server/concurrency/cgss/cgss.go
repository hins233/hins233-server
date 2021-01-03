package main

import (
	"bufio"
	"fmt"
	"os"
	"server/server/concurrency/cg"
	"server/server/concurrency/ipc"
	"strconv"
	"strings"
)

var centerClient *cg.CenterClient

func startCenterService() error {
	server := ipc.NewIpcServer(&cg.CenterServer{})
	client := ipc.NewIpcClient(server)
	centerClient = &cg.CenterClient{IpcClient: client}
	return nil
}

func Help(args []string) int {
	fmt.Println(`
		Commands:
			login <username><level><exp>
			logout <username>
			
	`)
	return 0
}

func Quit(args []string) int {
	return 1
}

func Logout(args []string) int {
	if len(args) != 2 {
		fmt.Println("USAGE: logout <username>")
	}
	centerClient.RemovePlayer(args[1])
	return 0
}

// login，把传入的参数解析成player对象。
func Login(args []string) int {
	if len(args) != 4 {
		fmt.Println("USAGE: login <username><level><exp>")
		return 0
	}
	player := cg.NewPlayer()
	player.Name = args[1]
	// 这里需要做格式转换，输入的值必须为integer类型
	player.Level, _ = strconv.Atoi(args[2])
	player.Exp, _ = strconv.Atoi(args[3])
	err := centerClient.AddPlayer(player)
	if err != nil {
		fmt.Println("failed adding player,", err)
	}
	return 0
}

func GetCommandHandlers() map[string]func(args []string) int {
	return map[string]func(args []string) int{
		"help":   Help,
		"quit":   Quit,
		"login":  Login,
		"logout": Logout,
	}
}

func main() {
	fmt.Println("Casual Game Server Solution")
	startCenterService()
	Help(nil)
	r := bufio.NewReader(os.Stdin)
	handlers := GetCommandHandlers()
	for {
		fmt.Print("Command>")
		b, _, _ := r.ReadLine()
		//r.ReadString('\n')
		line := string(b)
		tokens := strings.Split(line, ",")
		if handler, ok := handlers[tokens[0]]; ok {
			ret := handler(tokens)
			if ret != 0 {
				break
			}
		} else {
			fmt.Println("Unknown command:", tokens[0])
		}
	}
}
