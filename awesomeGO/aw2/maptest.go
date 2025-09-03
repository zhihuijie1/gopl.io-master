package aw2

// map类型的切片
func SliceMap(mapCap int) {
	slicemap := make([]map[string]int, mapCap)

	for i := 0; i < mapCap; i++ {
		slicemap[i] = make(map[string]int)
		str := "slicemap" + string(rune(i+'0'))
		slicemap[i][str] = i
	}

	for i := 0; i < mapCap; i++ {
		str := "slicemap" + string(rune(i+'0'))
		println("key是", str, "的value是", slicemap[i][str])
	}
}

// 值为切片类型的map
func MapSlice(mapCap int) {
	mapSlice := make(map[string][]string, mapCap)
	for i := 0; i < mapCap; i++ {
		println("第" + string(rune(i+'0')) + " 次循环")
		mapSliceKey := "mapSlice " + string(rune(i+'0'))

		mapSlice[mapSliceKey] = make([]string, i+1)

		for j := 0; j < i+1; j++ {
			mapSlice[mapSliceKey][j] = string(rune(j + '0'))
		}

		for j := 0; j < len(mapSlice[mapSliceKey]); j++ {
			println(mapSliceKey + " 的值是" + mapSlice[mapSliceKey][j])
		}
	}

}
