package jps
/**
 * 方向
 * */

type Direction struct {
	Id     int
	Desc   string
	Values []*Direction
	Divide []*Direction // 分散方向
}

/**
	八个方向，上下左右，上左 上右.....
 */
type UP struct {

}
