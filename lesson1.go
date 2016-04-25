package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {
	fmt.Println("my favorite number is ", rand.Intn(2), "\n")
	fmt.Printf("Now you have %g problems \n", math.Nextafter(2, 3))
	fmt.Println("Pi is ", math.Pi)
}
