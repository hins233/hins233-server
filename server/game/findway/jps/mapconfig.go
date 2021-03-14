package jps

type MapConfig struct {
	maps   []byte
	width  int
	height int
}

func InitMap(mapPosition [][]int) *MapConfig {
	width := len(mapPosition[0])
	height := len(mapPosition)
	maps := make([]byte, width*height)
	var i = 0
	for _, row := range mapPosition {
		for _, point := range row {
			maps[i] = byte(point)
			i++
		}
	}
	return &MapConfig{
		maps:   maps,
		width:  width,
		height: height,
	}
}

func (m *MapConfig) GetMap() []byte {
	return m.maps
}

func (m *MapConfig) GetWidth() int {
	return m.width
}

func (m *MapConfig) GetHeight() int {
	return m.height
}
