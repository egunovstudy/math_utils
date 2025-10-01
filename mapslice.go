package math_utils

func MapSlice(s []int, op func(int) int) {
	for _, n := range s {
		op(n)
	}
}
