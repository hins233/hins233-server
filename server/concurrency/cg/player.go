package cg

import "fmt"

type Player struct {
	Id    int
	Name  string
	Level int
	Exp   int
	Room  int

	mq chan *Message
}

func NewPlayer() *Player {
	m := make(chan *Message, 1024)
	player := &Player{0, "", 0, 0, 0, m}
	go func(p *Player) {
		for {
			select {
			case msg := <-p.mq:
				fmt.Println(p.Name, "Received message :", msg.Content)
			default:
				continue
			}
		}
	}(player)
	return player
}
