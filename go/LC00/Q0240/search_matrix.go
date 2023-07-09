package Q0240

func searchMatrix(matrix [][]int, target int) bool {
	i, j, l := 0, len(matrix[0])-1, len(matrix)
	for i < l && j >= 0 {
		if matrix[i][j] == target {
			return true
		} else if matrix[i][j] < target {
			i++
		} else if matrix[i][j] > target {
			j--
		}
	}
	return false
}
