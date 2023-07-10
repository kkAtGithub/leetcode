package binary_tree

// 基本的二叉树节点
type TreeNode struct {
	val   int
	left  *TreeNode
	right *TreeNode
}

// 后序遍历二叉树
func traverse(root *TreeNode) {
	if root != nil {
		// 前序位置
		traverse(root.left)
		// 中序位置
		traverse(root.right)
		// 后序位置
	}
}
