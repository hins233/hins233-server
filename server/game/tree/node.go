package tree

type Node interface {
	//AddChild(node Node) error
	RefreshAllNode(allNode map[int]Node) error
	Delete(allNode map[int]Node) error // 删掉当前节点
}

// FightNode 战力节点
type FightNode struct {
	Id             int // 节点id
	X              int // x 坐标
	Y              int // y 坐标
	FightPointSelf int // 节点自身的值
	FightPoint     int // 当前节点的总值
	Dad            int // 父节点，0为根节点
	Sons           map[int]*FightNode
}

//func (n *FightNode) AddChild(node Node) error {
//	newNode, ok := node.(*FightNode)
//	if !ok {
//		return fmt.Errorf("AddChild type assertion error node=%v", node)
//	}
//	n.Sons[newNode.Id] = newNode
//	newNode.Dad = n.Id
//	return nil
//}

func (n *FightNode) SetValue(value int) {
	n.FightPointSelf = value
}

func (n *FightNode) RefreshAllNode(allNode map[int]Node) error {
	temp := n
	ok := true
	for temp != nil && ok {
		temp.countValue()
		temp, ok = allNode[temp.Dad].(*FightNode)
	}
	return nil
}

func (n *FightNode) Delete(allNode map[int]Node) error {
	// 先统计要删的id列表
	var ids []int
	queue := []int{n.Id}
	for len(queue) > 0 {
		curr := queue[0]
		node, ok := allNode[curr].(*FightNode)
		if !ok {
			break
		}
		ids = append(ids, curr)
		for id := range node.Sons {
			queue = append(queue, id)
		}
		queue = queue[1:]
	}
	for _, delId := range ids {
		delete(allNode, delId)
	}
	return nil
}

func (n *FightNode) countValue() {
	n.FightPoint = 0
	for _, node := range n.Sons {
		n.FightPoint += node.FightPoint
	}
	n.FightPoint += n.FightPointSelf
}











