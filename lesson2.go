package main

import (
	"fmt"
)

func add(x int, y int) int {
	return x + y
}

//当两个或多个连续命名的参数是统一类型,则除最后一个要加类型之外，其他都可以省略
func addSecond(x, y, z int) int {
	return x + y + z
}

func main() {
	fmt.Println(add(3, 5))

	fmt.Println(addSecond(1, 4, 7))
}
