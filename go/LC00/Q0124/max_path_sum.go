package Q0124

import "math"

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var maxSum int

func maxPathSum(root *TreeNode) int {
	maxSum = math.MinInt
	backTrack(root)
	return maxSum
}

func backTrack(root *TreeNode) int {
	if root == nil {
		return 0
	}

	leftPathSum := max(0, backTrack(root.Left))
	rightPathSum := max(0, backTrack(root.Right))
	maxSum = max(maxSum, leftPathSum+rightPathSum+root.Val)
	return max(leftPathSum, rightPathSum) + root.Val
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
