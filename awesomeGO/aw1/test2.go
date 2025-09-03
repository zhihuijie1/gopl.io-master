package aw1

func CopyArray(arr *[5]int) {
	arr[0] = 10

	for i, v := range arr {
		println(i, "  ", v)
	}
}
