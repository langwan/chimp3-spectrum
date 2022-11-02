package sdk

type Frame struct {
	Name     string        `json:"name"`
	Filepath string        `json:"filepath"`
	IsPlay   bool          `json:"is_player"`
	Samples  *[2][]float64 `json:"samples"`
	Mode     int           `json:"mode"`
}

func SimplesSubConvert(simples *[][2]float64, depth int) [2][]float64 {
	var result [2][]float64
	l := depth
	if depth > len(*simples) {
		l = len(*simples)
	}
	result[0] = make([]float64, l)
	result[1] = make([]float64, l)
	for i, simple := range *simples {
		result[0][i] = simple[0]
		result[1][i] = simple[1]
	}
	return result
}
