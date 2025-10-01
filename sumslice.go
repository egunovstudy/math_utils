package math_utils

func SumSlice(s []int) (result int) {
	for _, n := range s {
		result += n
	}
	return result
}
