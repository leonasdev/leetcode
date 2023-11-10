package leetcode

func isValidSudoku(board [][]byte) bool {
	rowArr := [9][9]bool{}
	colArr := [9][9]bool{}
	squareArr := [9][9]bool{}

	for i, row := range board {
		for j, b := range row {
			if b == '.' {
				continue
			}
			v := int(b - '1')
			if rowArr[i][v] || colArr[j][v] || squareArr[i/3*3+j/3][v] {
				return false
			}
			rowArr[i][v], colArr[j][v], squareArr[i/3*3+j/3][v] = true, true, true
		}
	}

	return true
}
