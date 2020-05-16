package rotate_image

//
// question: https://leetcode.com/problems/rotate-image/
//

// Result
// Runtime: 0 ms, faster than 100.00% of Go online submissions for Rotate Image.
// Memory Usage: 2.3 MB, less than 100.00% of Go online submissions for Rotate Image.
func rotate(matrix [][]int) {
	newMatrix := make([][]int, len(matrix))
	y := 0
	for i := len(matrix) - 1; i > -1; i-- {
		array := make([]int, len(matrix[i]))
		x := i + y
		for j := 0; j < len(matrix[i]); j++ {
			array[j] = matrix[x][y]
			x--
		}
		newMatrix[y] = array
		y++
	}
	for i := 0; i < len(matrix); i++ {
		matrix[i] = newMatrix[i]
	}
}

// better answer because it is simpler than mine.
func rotate_v2(matrix [][]int) {
	m := len(matrix)
	n := len(matrix)

	for i := 0; i < m/2; i++ {

		for j := i; j < n-i-1; j++ {

			matrix[i][j], matrix[j][n-i-1], matrix[m-i-1][n-j-1], matrix[m-j-1][i] = matrix[m-j-1][i], matrix[i][j], matrix[j][n-i-1], matrix[m-i-1][n-j-1]

		}
	}
}
