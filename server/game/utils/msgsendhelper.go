package utils

import (
	"github.com/gobwas/ws"
	"github.com/gobwas/ws/wsutil"
	json "github.com/json-iterator/go"
	"log"
	"net"
)

func Send(conn net.Conn, resp map[string]interface{}, msgId int) error {
	resp["msgId"] = msgId
	res, err := json.Marshal(resp)
	log.Println("reply message :", string(res))
	err = wsutil.WriteServerMessage(conn, ws.OpText, res)
	if err != nil {
		log.Printf("write message error: %v", err)
		return err
	}
	return nil
}
