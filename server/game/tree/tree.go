package tree

import (
	"fmt"
	"math/rand"
)

const (
	BuilderGenerateFight = iota + 1
)

type Tree struct {
	Id      int
	root    Node
	allNode map[int]Node
}

func NewTree() *Tree {
	t := &Tree{
		Id:      0,
		root:    nil,
		allNode: make(map[int]Node),
	}
	err := t.buildNode(BuilderGenerateFight, nil)
	if err != nil {
		panic(err)
	}
	return t
}

func (t *Tree) AddChild(nodeId int) error {
	node, ok := t.allNode[nodeId]
	if !ok {
		return fmt.Errorf("不存在此node")
	}
	return t.buildNode(BuilderGenerateFight, node)
}

func (t *Tree) RemoveNode(nodeId int) error {
	delNode, ok := t.allNode[nodeId]
	if !ok {
		return fmt.Errorf("del node not exit")
	}
	if delNode == t.root {
		return fmt.Errorf("can't del root node")
	}
	return delNode.Delete(t.allNode)
}

func (t *Tree) ChangePos(nodeId, x, y int) error {
	node, ok := t.allNode[nodeId].(*FightNode)
	if !ok {
		return fmt.Errorf("不存在此node")
	}
	node.X = x
	node.Y = y
	return nil
}

func (t *Tree) ToMap() map[string]interface{} {
	toM := make(map[string]interface{})
	toM["allNode"] = t.allNode
	return toM
}

var buildNodeRegister = make(map[int]func(t *Tree, dad Node) error)

func (t *Tree) buildNode(module int, dad Node) error {
	builder, ok := buildNodeRegister[module]
	if ok {
		return builder(t, dad)

	}
	return fmt.Errorf("build node factory have not module:%d", module)
}

func (t *Tree) getId() int {
	id := t.Id
	t.Id++
	return id
}

func buildFightNode(t *Tree, dad Node) error {
	node := &FightNode{
		Id:             t.getId(),
		FightPointSelf: rand.Intn(100) + 1,
		Sons:           map[int]*FightNode{},
	}
	t.allNode[node.Id] = node
	if dad == nil {
		node.Dad = -1
		node.X = 400
		node.Y = 200
		node.FightPoint = node.FightPointSelf
		return nil
	}
	node.X = dad.(*FightNode).X + 20
	node.Y = dad.(*FightNode).Y + 20
	node.Dad = dad.(*FightNode).Id
	return node.RefreshAllNode(t.allNode)
}

func registerBuildNodeFunc(module int, fn func(t *Tree, dad Node) error) {
	buildNodeRegister[module] = fn
}

func init() {
	registerBuildNodeFunc(BuilderGenerateFight, buildFightNode)
}
