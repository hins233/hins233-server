package jps

import "server/server/game/findway"

type JPS struct {
	MapCfg    findway.MapConfig
	openList  []*JumpPoint
	closeSet  map[int]*JumpPoint
	MapPoints []*Point
}

// 跳点
type JumpPoint struct {
	Weight int // 权重
	Index  int
	Parent *JumpPoint
}

type Point struct {
	X         int
	Y         int
	Neighbors []*Neighbor
}

type Neighbor struct {
	Index  int
	Common []int // 共同方向的点
}
