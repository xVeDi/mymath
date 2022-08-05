package mymath

// Add 2 int values
func Add(a, b int) int {
	return a + b
}

// AddArr Add 2 int slices
func AddArr(a, b []int) []int {
	length := len(a)
	if length-len(b) < 0 {
		length = len(b)
	}
	c := make([]int, length)

	for i := 0; i < length; i++ {
		aa := a[i]
		//if ok != nil {
		//	aa = 0
		//}
		bb := b[i]
		//if ok != nil {
		//	bb = 0
		//}

		c[i] = aa + bb
	}

	return c
}
