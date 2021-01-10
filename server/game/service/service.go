package service

import (
	"errors"
	"fmt"
	"log"
	"net"
	"server/server/game/tree"
)

var (
	serviceMap = make(map[int][]GameService)
)

type GameService interface {
	Service(param map[string]interface{}, conn net.Conn) error
}

func Action(gameId, msgId int, data map[string]interface{}, conn net.Conn) error {
	services, ok := serviceMap[gameId]
	if !ok {
		return errors.New(fmt.Sprintf("service map not have gameId: %d", gameId))
	}
	if len(services) <= msgId {
		log.Panicf("msgId invalid len=%d: %d", len(services), msgId)
		return errors.New("msgId invalid")
	}
	return services[msgId].Service(data, conn)
}

func RegisterService(gameId int, service GameService) {
	serviceMap[gameId] = append(serviceMap[gameId], service)
}

func init() {
	RegisterService(1, &tree.TestController{})
	RegisterService(1, &tree.StartController{})
	RegisterService(1, &tree.RemoveController{})
	RegisterService(1, &tree.AddController{})
	RegisterService(1, &tree.ChangePosController{})
}
