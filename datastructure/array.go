package datastructure

import "playground/utils"

func showArray() {
	var arr1 [3]int                   // 长度为3的整数数组，元素初始化为零值
	var arr2 [3]int = [3]int{1, 2, 3} // 长度为3的整数数组，显式初始化元素
	arr3 := [...]int{1, 2, 3}         // Go会根据初始化值的数量推导数组大小
	arr4 := [5]int{2: 10, 4: 20}      // 索引为2和4的元素分别初始化为10和20，其他元素为零值

	var matrix1 [2][3]int                      // 二维数组，2行3列，元素初始化为零值
	matrix2 := [2][3]int{{1, 2, 3}, {4, 5, 6}} // 二维数组，显式初始化元素

	utils.Print(
		arr1,
		arr2,
		arr3,
		arr4,
		matrix1,
		matrix2,
	)
}
