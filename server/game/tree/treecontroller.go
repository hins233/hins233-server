package tree

import (
	"errors"
	"log"
	"net"
	"server/server/game/utils"
)

var (
	treeMap = make(map[string]*Tree)
)

type StartController struct {
}

func (c *StartController) Service(param map[string]interface{}, conn net.Conn) error {
	t := NewTree()
	treeMap[conn.RemoteAddr().String()] = t
	resp := t.ToMap()
	return utils.Send(conn, resp, 1)
}

type RemoveController struct {
}

func (c *RemoveController) Service(param map[string]interface{}, conn net.Conn) error {
	id, ok := param["id"].(float64)
	if !ok {
		return errors.New("param not valid")
	}
	t, ok := treeMap[conn.RemoteAddr().String()]
	if !ok {
		log.Panicf("tree map not find key %s", conn.RemoteAddr().String())
		return errors.New("tree map err")
	}
	t.RemoveNode(int(id))
	return utils.Send(conn, t.ToMap(), 1)
}

type AddController struct {
}

func (c *AddController) Service(param map[string]interface{}, conn net.Conn) error {
	id, ok := param["id"].(float64)
	if !ok {
		return errors.New("param not valid")
	}
	t, ok := treeMap[conn.RemoteAddr().String()]
	if !ok {
		log.Panicf("tree map not find key %s", conn.RemoteAddr().String())
		return errors.New("tree map err")
	}
	t.AddChild(int(id))
	return utils.Send(conn, t.ToMap(), 1)
}

type ChangePosController struct {
}

func (c *ChangePosController) Service(param map[string]interface{}, conn net.Conn) error {
	id, ok1 := param["id"].(float64)
	x, ok2 := param["x"].(float64)
	y, ok3 := param["y"].(float64)
	if !(ok1 && ok2 && ok3) {
		return errors.New("param not valid")
	}
	t, ok := treeMap[conn.RemoteAddr().String()]
	if !ok {
		log.Panicf("tree map not find key %s", conn.RemoteAddr().String())
		return errors.New("tree map err")
	}
	t.ChangePos(int(id), int(x), int(y))
	return utils.Send(conn, t.ToMap(), 1)
}

type TestController struct {
}

func (c *TestController) Service(param map[string]interface{}, conn net.Conn) error {
	resp := make(map[string]interface{})
	resp["1"] = 1
	return utils.Send(conn, resp, 0)
}

//func init() {
//	service.RegisterService(1, &TestController{})
//	service.RegisterService(1, &StartController{})
//	service.RegisterService(1, &RemoveController{})
//	service.RegisterService(1, &AddController{})
//	service.RegisterService(1, &ChangePosController{})
//}
