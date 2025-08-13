package main

import (
	"math"
	"sort"
)

func Hammer(n int) uint {

	arr := make([]uint, 0)
	var limit uint = math.MaxUint
	size := 64

	for i := 0; i < size; i++ {
		a := PowerOf(2, i)
		for j := 0; j < size; j++ {
			b := PowerOf(3, j)
			if b > 0 && a >= limit/b {
				break
			}

			ab := a * b
			for k := 0; k < size; k++ {
				c := PowerOf(5, k)
				if c > 0 && ab >= limit/c {
					break
				}
				arr = append(arr, ab*c)
			}
		}
	}

	sort.Slice(arr, func(i, j int) bool {
		return arr[i] < arr[j]
	})

	return arr[n-1]
}

func PowerOf(m uint, n int) uint {
	var result uint = 1
	for i := 0; i < n; i++ {
		result *= m
	}
	return result
}

func dotest(n int, exp uint) {
	Expect(Hammer(n), exp)
}

func TestFunc() {
	dotest(1, 1)
	dotest(2, 2)
	dotest(3, 3)
	dotest(4, 4)
	dotest(5, 5)
	dotest(6, 6)
	dotest(7, 8)
	dotest(8, 9)
	dotest(9, 10)
	dotest(10, 12)
	dotest(11, 15)
	dotest(12, 16)
	dotest(13, 18)
	dotest(14, 20)
	dotest(15, 24)
	dotest(16, 25)
	dotest(17, 27)
	dotest(18, 30)
	dotest(19, 32)
}
