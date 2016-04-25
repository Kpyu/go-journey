package main

import (
	"fmt"
)

/** 函数可以返回任意数量的返回值**/

func swap(a, b, c string) (string, string, string) {
	return c, b, a
}

func main() {
	a, b, c := swap("world", "my", "hello")

	fmt.Println(a, b, c)
}
