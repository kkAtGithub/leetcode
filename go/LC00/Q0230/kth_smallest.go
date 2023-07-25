package Q0230

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

var kthNum int
var current int

func kthSmallest(root *TreeNode, k int) int {
	kthNum = -1
	current = 0
	traverse(root, k)
	return kthNum
}

func traverse(root *TreeNode, k int) {
	if root == nil {
		return
	}
	traverse(root.Left, k)
	current++
	if current == k {
		kthNum = root.Val
		return
	}
	traverse(root.Right, k)
}
