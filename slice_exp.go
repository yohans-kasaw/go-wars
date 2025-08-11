package main

import "fmt"

func mainMat() {
	s := [][]int{{1}, {2}}
	modify(s)

	fmt.Println(s)
}

func modify(s [][]int) {
	s[0][0] = 1000
}

func TestFunc() {
	mainMat()
}
