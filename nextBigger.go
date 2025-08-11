package main

func NextBigger(n int) int {
	digits := int_to_arr(n)

	// Find the rightmost digit that is smaller than the digit next to it
	pivot := -1
	for i := len(digits) - 2; i >= 0; i-- {
		if digits[i] < digits[i+1] {
			pivot = i
			break
		}
	}

	// If no such digit is found, the number is the largest permutation
	if pivot == -1 {
		return -1
	}

	// Find the smallest digit on right side of pivot that is greater than pivot
	successor := -1
	for i := len(digits) - 1; i > pivot; i-- {
		if digits[i] > digits[pivot] {
			successor = i
			break
		}
	}

	// Swap the pivot and successor
	digits[pivot], digits[successor] = digits[successor], digits[pivot]

	// Reverse the suffix starting at pivot+1
	left, right := pivot+1, len(digits)-1
	for left < right {
		digits[left], digits[right] = digits[right], digits[left]
		left++
		right--
	}

	return arr_to_int(digits)
}

func int_to_arr(n int) []int {
	arr := []int{}
	for ; n > 0; n /= 10 {
		arr = append(arr, n%10)
	}
	return arr
}

func arr_to_int(arr []int) int {
	n := 0
	for _, digit := range arr {
		n = n*10 + digit
	}
	return n
}

func TestFunc() {
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
