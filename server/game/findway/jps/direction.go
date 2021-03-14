package jps

/**
 * 方向
 * */
type DirectionInterface interface {
	// 取出这个方向的点，（x,y）代表点，(width,height)代表横纵单位长度。为啥不直接设死呢，因为宽和高是由前端设置的
	GetPoint(x, y, width, height int) int

	// 获取相同的点
	GetCommon(x, y, width, height int) []int

	// 取反方向
	GetReverse() DirectionInterface

	// 获取可选方向
	GetDivide() []*Direction

	GetId() int

	SetId(id int)

	SetDesc(desc string)

	// 设置方向
	SetDivide(divide []*Direction)
}

type Direction struct {
	Id     int          // 方向的id，枚举值
	Desc   string       // 方向的描述
	Values []*Direction // 当前点有哪些方向。
	Divide []*Direction // 分散方向
}

func (d *Direction) GetPoint(x, y, width, height int) int {
	return 0
}

func (d *Direction) GetCommon(x, y, width, height int) []int {
	return nil
}

func (d *Direction) GetReverse() DirectionInterface {
	return nil
}

func (d *Direction) GetDivide() []*Direction {
	return d.Divide
}

func (d *Direction) SetDesc(desc string) {
	d.Desc = desc
}

func (d *Direction) GetId() int {
	return d.Id
}

func (d *Direction) SetId(id int) {
	d.Id = id
}

func (d *Direction) SetDivide(divide []*Direction) {
	d.Divide = divide
}

/**
八个方向，上下左右，上左 上右.....
*/
/**
 * 方向
 * */
type DirectionFactory func(id int, desc string) DirectionInterface

var (
	UP            = Generate(0, "上")
	RIGHT         = Generate(1, "右")
	DOWN          = Generate(2, "下")
	LEFT          = Generate(3, "左")
	UR            = Generate(4, "上右")
	DR            = Generate(5, "下右")
	DL            = Generate(6, "下左")
	UL            = Generate(7, "上左")
	DirectionList = [8]DirectionInterface{}
)

func Generate(id int, desc string, divide ...*Direction) DirectionInterface {
	var d DirectionInterface
	switch id {
	case 0:
		d = &up{}
	case 1:
		d = &right{}
	case 2:
		d = &down{}
	case 3:
		d = &left{}
	case 4:
		d = &ur{}
	case 5:
		d = &dr{}
	case 6:
		d = &dl{}
	case 7:
		d = &ul{}
	default:
		d = &Direction{}
	}
	d.SetId(id)
	d.SetDesc(desc)
	d.SetDivide(divide)
	DirectionList[id] = d
	return d
}

// todo 查一下有没有直接在func中实现接口的写法，感觉写在函数外不像是枚举了。
type up struct {
	Direction
}

func (d *up) GetPoint(x, y, width, height int) int {
	if y <= 0 {
		return -1
	}
	return (y-1)*width + x
}

func (d *up) GetCommon(x, y, width, height int) []int {
	return []int{LEFT.GetPoint(x, y, width, height),
		RIGHT.GetPoint(x, y, width, height),
		UR.GetPoint(x, y, width, height),
		UL.GetPoint(x, y, width, height)}
}

func (d *up) GetReverse() DirectionInterface {
	return DOWN
}

type right struct {
	Direction
}

func (d *right) GetPoint(x, y, width, height int) int {
	if x >= width-1 {
		return -1
	}
	return y*width + x + 1
}

func (d *right) GetCommon(x, y, width, height int) []int {
	return []int{UP.GetPoint(x, y, width, height),
		DOWN.GetPoint(x, y, width, height),
		UR.GetPoint(x, y, width, height),
		DR.GetPoint(x, y, width, height)}
}

func (d *right) GetReverse() DirectionInterface {
	return LEFT
}

type down struct {
	Direction
}

func (d *down) GetPoint(x, y, width, height int) int {
	if y >= height-1 {
		return -1
	}
	return (y+1)*width + x
}

func (d *down) GetCommon(x, y, width, height int) []int {
	return []int{RIGHT.GetPoint(x, y, width, height),
		LEFT.GetPoint(x, y, width, height),
		DL.GetPoint(x, y, width, height),
		DR.GetPoint(x, y, width, height)}
}

func (d *down) GetReverse() DirectionInterface {
	return UP
}

type left struct {
	Direction
}

func (d *left) GetPoint(x, y, width, height int) int {
	if x <= 0 {
		return -1
	}
	return y*width + x - 1
}

func (d *left) GetCommon(x, y, width, height int) []int {
	return []int{UP.GetPoint(x, y, width, height),
		DOWN.GetPoint(x, y, width, height),
		UL.GetPoint(x, y, width, height),
		DL.GetPoint(x, y, width, height)}
}

func (d *left) GetReverse() DirectionInterface {
	return RIGHT
}

type ur struct {
	Direction
}

func (d *ur) GetPoint(x, y, width, height int) int {
	if y <= 0 || x >= width-1 {
		return -1
	}
	return (y-1)*width + x + 1
}

func (d *ur) GetCommon(x, y, width, height int) []int {
	return []int{UP.GetPoint(x, y, width, height),
		RIGHT.GetPoint(x, y, width, height)}
}

func (d *ur) GetReverse() DirectionInterface {
	return DL
}

type dr struct {
	Direction
}

func (d *dr) GetPoint(x, y, width, height int) int {
	if y >= height-1 || x >= width-1 {
		return -1
	}
	return (y+1)*width + x + 1
}

func (d *dr) GetCommon(x, y, width, height int) []int {
	return []int{RIGHT.GetPoint(x, y, width, height),
		DOWN.GetPoint(x, y, width, height)}
}

func (d *dr) GetReverse() DirectionInterface {
	return UL
}

type dl struct {
	Direction
}

func (d *dl) GetPoint(x, y, width, height int) int {
	if y >= height-1 || x <= 0 {
		return -1
	}
	return (y+1)*width + x - 1
}

func (d *dl) GetCommon(x, y, width, height int) []int {
	return []int{LEFT.GetPoint(x, y, width, height),
		DOWN.GetPoint(x, y, width, height)}
}

func (d *dl) GetReverse() DirectionInterface {
	return UR
}

type ul struct {
	Direction
}

func (d *ul) GetPoint(x, y, width, height int) int {
	if y <= 0 || x <= 0 {
		return -1
	}
	return (y-1)*width + x - 1
}

func (d *ul) GetCommon(x, y, width, height int) []int {
	return []int{UP.GetPoint(x, y, width, height),
		LEFT.GetPoint(x, y, width, height)}
}

func (d *ul) GetReverse() DirectionInterface {
	return DR
}
