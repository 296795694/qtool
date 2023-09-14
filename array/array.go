package array

// InArrayString 判断是否在数组内
func InArrayString(items []string, item string) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// InArrayInt 判断是否在数组内
func InArrayInt(items []int, item int) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// InArrayUint 判断是否在数组内
func InArrayUint(items []uint, item uint) bool {
	for _, eachItem := range items {
		if eachItem == item {
			return true
		}
	}
	return false
}

// Combine 数组排列组合
func Combine(sets ...[]string) [][]string {
	lens := func(i int) int { return len(sets[i]) }
	var array [][]string
	for ix := make([]int, len(sets)); ix[0] < lens(0); nextIndex(ix, lens) {
		var r []string
		for j, k := range ix {
			r = append(r, sets[j][k])
		}
		array = append(array, r)
	}
	return array
}

// nextIndex 获取下一级
func nextIndex(ix []int, lens func(i int) int) {
	for j := len(ix) - 1; j >= 0; j-- {
		ix[j]++
		if j == 0 || ix[j] < lens(j) {
			return
		}
		ix[j] = 0
	}
}

// Reverse 反转
func Reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
