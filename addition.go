package mymath

func Add(a, b int) int {
	return a + b
}

func AddArr(a, b []int) []int {
	length := len(a)
	if length-len(b) > 0 {
		length = len(b)
	}
	c := make([]int, length)

	for i := 0; i < length; i++ {
		c[i] = a[i] + b[i]
	}

	return c
}
