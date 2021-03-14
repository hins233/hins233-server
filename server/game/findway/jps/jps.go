package jps

import "math"

type JPS struct {
	MapCfg    *MapConfig
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
	X         int         `json:"x"`
	Y         int         `json:"y"`
	Neighbors []*Neighbor `json:"neighbors"`
}

type Neighbor struct {
	Index  int   `json:"index"`
	Common []int `json:"common"` // 共同方向的点
}

func InitJPS(mapCfg *MapConfig) *JPS {
	MapPoints := make([]*Point, len(mapCfg.maps))
	i := 0
	for y := 0; y < mapCfg.GetHeight(); y++ {
		for x := 0; x < mapCfg.GetWidth(); x++ {
			MapPoints[i] = generatePoint(x, y, mapCfg.GetWidth(), mapCfg.GetHeight())
			i++
		}
	}
	return &JPS{
		MapCfg:    mapCfg,
		openList:  nil,
		closeSet:  make(map[int]*JumpPoint),
		MapPoints: MapPoints,
	}
}

func generatePoint(x, y, width, height int) *Point {
	neighbors := make([]*Neighbor, 8)
	for i, direct := range DirectionList {
		neighbors[i] = &Neighbor{
			Index:  direct.GetPoint(x, y, width, height),
			Common: direct.GetCommon(x, y, width, height),
		}
	}
	return &Point{
		X:         x,
		Y:         y,
		Neighbors: neighbors,
	}
}

func (j *JPS) FindPath(startX, startY, endX, endY int) []*Point {
	sI := startY*j.MapCfg.GetWidth() + startX
	eI := endY*j.MapCfg.GetWidth() + endX

	start := generateJumpPoint(sI, 0, nil)
	j.openList = append(j.openList, start)
	j.closeSet[start.Index] = start
	var end *JumpPoint
	for len(j.openList) > 0 {
		jumpPoint := j.openList[0]
		j.openList = j.openList[1:]
		end = j.horizontalFind(jumpPoint, eI)
		if end != nil {
			break
		}
		end = j.obliqueFind(jumpPoint, eI)
		if end != nil {
			break
		}
	}
	if end == nil {
		return nil
	}
	var route []*Point
	for end != nil {
		route = append([]*Point{j.MapPoints[end.Index]}, route...)
		end = end.Parent
	}
	return route

}

// 水平找
func (j *JPS) horizontalFind(jumpPoint *JumpPoint, eI int) *JumpPoint {
	point := j.MapPoints[jumpPoint.Index]
	for i := 0; i < 4; i++ {
		neighbor := point.Neighbors[i]
		for neighbor.Index != -1 && j.MapCfg.GetMap()[neighbor.Index] != 1 {
			//提前找到终点
			if neighbor.Index == eI {
				return generateJumpPoint(eI, 0, jumpPoint)
			}
			j.findMustNeighbors(jumpPoint, neighbor.Index, DirectionList[i])
			neighborsPoint := j.MapPoints[neighbor.Index]
			neighbor = neighborsPoint.Neighbors[i]
		}
	}
	return nil
}

/**
 * 斜方向寻找
 * */
func (j *JPS) obliqueFind(jumpPoint *JumpPoint, eI int) *JumpPoint {
	point := j.MapPoints[jumpPoint.Index]
	for i := 4; i < 8; i++ {
		direction := DirectionList[i]
		neighbor := point.Neighbors[i]
		var distance float64 = 1.4
		for neighbor.Index != -1 && j.MapCfg.GetMap()[neighbor.Index] != 1 {
			if neighbor.Index == eI {
				return generateJumpPoint(eI, 0, jumpPoint)
			}
			//遍历分向
			nowPoint := j.MapPoints[neighbor.Index]
			//拐点
			now := generateJumpPoint(neighbor.Index, jumpPoint.Weight+int(math.Ceil(distance)), jumpPoint)
			flag := false
			for _, divide := range direction.GetDivide() {
				divideNeighbor := nowPoint.Neighbors[divide.GetId()]
				for divideNeighbor.Index != -1 && j.MapCfg.GetMap()[divideNeighbor.Index] != 1 {
					if divideNeighbor.Index == eI {
						return generateJumpPoint(eI, 0, now)
					}
					if j.findMustNeighbors(now, divideNeighbor.Index, divide) {
						flag = true
					}
					nowPoint = j.MapPoints[divideNeighbor.Index]
					divideNeighbor = nowPoint.Neighbors[divide.GetId()]
				}
				nowPoint = j.MapPoints[neighbor.Index]
			}
			if flag {
				j.closeSet[now.Index] = now
			}
			j.findMustNeighbors(jumpPoint, neighbor.Index, DirectionList[i])
			neighbor = nowPoint.Neighbors[i]
			distance += 1.4
		}
	}
	return nil
}

/**
 * 寻找强迫邻居 和 跳点
 * */
func (j *JPS) findMustNeighbors(jumpPoint *JumpPoint, nI int, direction DirectionInterface) bool {
	beforeJumpPoint := j.closeSet[nI]
	if beforeJumpPoint != nil {
		nowWeight := j.getDistance(nI, jumpPoint.Index) + jumpPoint.Weight
		if beforeJumpPoint.Weight > nowWeight {
			beforeJumpPoint.Parent = jumpPoint
			beforeJumpPoint.Weight = nowWeight
			j.openList = append(j.openList, beforeJumpPoint)
			return true
		}
		return false
	}
	nP := j.MapPoints[nI]
	for i := 0; i < len(nP.Neighbors); i++ {
		if i == direction.GetId() || i == direction.GetReverse().GetId() {
			continue
		}

		neighbor := nP.Neighbors[i]
		if neighbor.Index == -1 {
			continue
		}
		//有墙壁
		if j.MapCfg.GetMap()[neighbor.Index] == 1 {
			for _, commonIndex := range neighbor.Common {
				if commonIndex != -1 && j.MapCfg.GetMap()[commonIndex] == 0 { //与墙壁的公共点有非阻挡
					jpsJumpPoint := generateJumpPoint(nI, j.getDistance(nI, jumpPoint.Index)+jumpPoint.Weight, jumpPoint)
					j.openList = append(j.openList, jpsJumpPoint)
					j.closeSet[nI] = jpsJumpPoint
					return true
				}
			}
		}
	}
	return false
}

func (j *JPS) getDistance(start, end int) int {
	startP := j.MapPoints[start]
	endP := j.MapPoints[end]
	return int(math.Sqrt(math.Pow(float64(startP.X-endP.X), 2) + math.Pow(float64(startP.Y-endP.Y), 2)))
}

func generateJumpPoint(index, weight int, parent *JumpPoint) *JumpPoint {
	return &JumpPoint{
		Index:  index,
		Weight: weight,
		Parent: parent,
	}
}
