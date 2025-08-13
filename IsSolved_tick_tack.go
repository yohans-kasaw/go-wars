package main

func IsSolved(board [3][3]int) int {

	n := len(board)
	// check rows
	for _, row := range board {
		if checkArr(row) == true {
			return row[0]
		}
	}

	// check cols
	for j := range board {
		var col [3]int
		for i := range board {
			col[i] = board[i][j]
		}

		if checkArr(col) == true {
			return col[0]
		}
	}

	var right_accros [3]int
	var left_accros [3]int
	for i := range board {
		right_accros[i] = board[i][i]
		left_accros[i] = board[i][n-i-1]
	}

	if checkArr(right_accros) == true {
		return right_accros[0]
	}

	if checkArr(left_accros) == true {
		return left_accros[0]
	}

	if countZeros(board) == 0 {
		return 0
	}

	return -1
}

func countZeros(board [3][3]int) int {
	res := 0
	for _, row := range board {
		for _, v := range row {

			if v == 0 {

				res += 1
			}
		}
	}

	return res
}

func checkArr(arr [3]int) bool {
	for _, v := range arr {
		if v == 0 || v != arr[0] {
			return false
		}
	}
	return true
}
