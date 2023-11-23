package datastructure

import (
	"playground/utils"
	"sort"
)

func showSlice() {
	// defintion and init
	arr1 := [5]int{0, 1, 2, 3, 4}

	s1 := arr1[1:3]
	s2 := []int{1, 2, 3}
	s3 := make([]int, 3)

	s4 := append(s1, 5, 6)
	utils.Print(
		arr1,
		s1,
		s2,
		s3,
		s4,
		arr1,
	)

	// append
	s5 := append(s1, 5, 6, 7)
	s6 := append(s1, s5...)

	indexToDelete := 6
	s7 := append(s6[:indexToDelete], s6[indexToDelete+1:]...)

	utils.Print(
		s1,
		s4,
		s5,
		s6,
		s7,
	)

	// copy
	source := []int{1, 2, 3}
	target := make([]int, 2)
	copyCount := copy(target, source)
	utils.Print(
		target,
		copyCount,
	)

	// sort
	sort.Slice(s6, func(i, j int) bool {
		return s6[i] < s6[j]
	})

	utils.Print(
		s6,
	)

	// search
	k := 1
	n := len(s6)
	index := sort.Search(n, func(i int) bool {
		return s6[i] > k
	})

	utils.Print(index)

	s8 := make([]int, 0, 10)
	s8 = append(s8, 1, 2, 3)
	utils.Print(s8)

	var arr []int64
	arr = append(arr, 1, 2, 3, 4, 5)
	utils.Print(arr)
}
