package cg

import (
	"encoding/json"
	"errors"
	"server/server/concurrency/ipc"
)

type CenterClient struct {
	*ipc.IpcClient
}

func (client *CenterClient) AddPlayer(player *Player) error {
	b, err := json.Marshal(player)
	if err != nil {
		return err
	}
	resp, err := client.Call("addplayer", string(b))
	if err == nil && resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}

func (client *CenterClient) RemovePlayer(username string) error {
	resp, err := client.Call("removeplayer", username)
	if err == nil && resp.Code == "200" {
		return nil
	}
	return errors.New(resp.Code)
}
