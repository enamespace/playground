package datastructure

import (
	"fmt"
	"playground/utils"
)

func showHashmap() {
	// defintion and init
	var hashmap1 map[string]string // nil，需要使用make来初始化
	hashmap2 := map[string]string{}
	hashmap3 := make(map[string]string, 3)
	utils.Print(
		hashmap1,
		hashmap2,
		hashmap3,
	)

	// set, get, delete
	hashmap2["a"] = "va"

	utils.Print(
		hashmap2["a"],
		hashmap2["b"],
	)

	v, ok := hashmap2["a"]
	utils.Print(
		v,
		ok,
	)

	delete(hashmap2, "a")
	for k, v := range hashmap1 {
		fmt.Print(k, v)
	}
}
