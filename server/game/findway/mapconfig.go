package findway

type MapConfig struct {
	maps   []byte
	width  int
	height int
}

func InitMap(mapPosition [][]int) *MapConfig {
	width := len(mapPosition[0])
	height := len(mapPosition)
	maps := make([]byte, 0)
	for _, row := range mapPosition {
		for _, point := range row {
			maps = append(maps, byte(point))
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
