package main


func Determinant(matrix [][]int) int {

	if len(matrix) == 1 {
		return matrix[0][0]
	}

	res := 0
	for j := 0; j < len(matrix); j++ {
		sign := 1
		if j%2 == 1 {
			sign = -1
		}
		a := matrix[0][j]
		minor := makeMinor(&matrix, j)
		res += sign * a * Determinant(minor)
	}

	return res
}

func makeMinor(matrix *[][]int, j int) [][]int {

	n := len(*matrix)

	minor := make([][]int, 0, n-1)

	for i := 1; i < len(*matrix); i++ {
		row := (*matrix)[i]
		minor_row := make([]int, 0, n-1)

		minor_row = append(minor_row, row[:j]...)
		minor_row = append(minor_row, row[j+1:]...)

		minor = append(minor, minor_row)
	}

	return minor
}

func TestFunc() {
	Expect(Determinant([][]int{{1}}))
	Expect(Determinant([][]int{{1, 3}, {2, 5}}))
	Expect(Determinant([][]int{{2, 5, 3}, {1, -2, -1}, {1, 3, 4}}))
}
