package main

import (
	"fmt"
)

// func dotest(t, k int, ls []int, expected int) {
// 	fmt.Println(ChooseBestSum(t, k, ls), expected, t)
// }

func Expect[T comparable](res T, expected T) {
	fmt.Println(res == expected, res, "expected -> ", expected)
}

func main() {
	TestFunc()
}
