package datastructure

import (
	"fmt"
	"playground/utils"
	"strings"
)

func showString() {

	var s1 string
	s2 := "hello world"
	s3 := "你好，世界abc"
	utils.Print(
		s1,
		s2,
		s3,
	)

	utils.Print(
		s2[2:],
		s3[2:],
	)

	rs := []rune(s3)
	for i := 0; i < len(rs); i++ {
		fmt.Print(string(rs[i]), " ")
	}

	fmt.Println()
	bs := []byte(s2)
	for i := 0; i < len(bs); i++ {
		fmt.Print(string(bs[i]), " ")
	}

	utils.Print(
		strings.Contains("Hello, World!", "World"),
		strings.HasPrefix("Hello, World!", "Hello"),
		strings.HasSuffix("Hello, World!", "World!"),
	)

}
