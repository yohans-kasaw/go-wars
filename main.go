package main

import (
	"fmt"
	"sort"
)

func NextBigger(n int) int {
	arr := int_to_arr(n)

	first_dec := -1
	for i := 1; i < len(arr); i++ {
		if arr[i] < arr[i-1] {
			first_dec = i
			break
		}
	}

	if first_dec == -1 {
		return -1
	}

	for j := 0; j < first_dec; j++ {
		if arr[j] > arr[first_dec] {
			arr[first_dec], arr[j] = arr[j], arr[first_dec]
			sort.Sort(sort.Reverse(sort.IntSlice(arr[:first_dec])))
			res := arr_to_rev_int(&arr)
			return res
		}
	}
	return -1
}

func int_to_arr(n int) []int {
	arr := []int{}
	for ; n > 0; n /= 10 {
		arr = append(arr, n%10)
	}
	return arr
}

func arr_to_rev_int(arr *[]int) int {
	n := 0
	for i := len(*arr) - 1; i >= 0; i-- {
		n *= 10
		n += (*arr)[i]
	}
	return n
}

// func dotest(t, k int, ls []int, expected int) {
// 	fmt.Println(ChooseBestSum(t, k, ls), expected, t)
// }

func Expect(res int) {
	fmt.Println(res)
}

func main() {
	example_tests := [...][]int{
		{8, -1},
		{12, 21},
		{513, 531},
		{2017, 2071},
		{414, 441},
		{144, 414},
		{1234567890, -1},
		{9876543210, -1},
		{59884848459853, -1},
	}
	for _, v := range example_tests {
		input := v[0]
		Expect(NextBigger(input))
	}
}
