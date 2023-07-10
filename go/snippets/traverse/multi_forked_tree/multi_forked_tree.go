package multi_forked_tree

// 基本的 N 叉树节点
type TreeNode struct {
	val      int
	children []*TreeNode
}

func traverse(root *TreeNode) {
	for _, child := range root.children {
		traverse(child)
	}
}
