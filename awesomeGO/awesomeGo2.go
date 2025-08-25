package awesomeGO // 保持包声明不变

var arr [2][3]int

func Arrays() [2][3]int {
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			arr[i][j] = i * j
		}
	}
	return arr
}
