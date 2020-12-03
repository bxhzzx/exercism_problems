package pascal

func Triangle(size int) [][]int {
	res := make([][]int, size)
	if size == 1 {
		res[0] = []int{1}
		return res
	}
	res[0] = []int{1}
	res[1] = []int{1, 1}

	for i := 2; i < size; i++ {
		tmpSlice := make([]int, i+1)
		for i2 := range tmpSlice {
			if i2 == 0 || i2 == len(tmpSlice)-1 {
				tmpSlice[i2] = 1
				continue
			}
			tmpSlice[i2] = res[i-1][i2-1] + res[i-1][i2]
		}
		res[i] = tmpSlice
	}
	return res
}
