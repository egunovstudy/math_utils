package math_utils

func FoldSlice(s []int, op func(int, int) int, init int) (result int) {
	result = op(init, s[0])
	for i := 1; i < len(s); i++ {
		result = op(result, s[i])
	}
	return result
}
