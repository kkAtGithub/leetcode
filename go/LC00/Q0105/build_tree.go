package Q0105

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	return build(preorder, inorder, 0, len(preorder)-1, 0, len(inorder)-1)
}

func build(preorder, inorder []int, preStart, preEnd, inStart, inEnd int) *TreeNode {
	if preStart > preEnd || inStart > inEnd {
		return nil
	}
	rootVal := preorder[preStart]
	rootIndex := 0
	for i := inStart; i <= inEnd; i++ {
		if inorder[i] == rootVal {
			rootIndex = i
			break
		}
	}

	leftSize := rootIndex - inStart

	return &TreeNode{
		Val:   rootVal,
		Left:  build(preorder, inorder, preStart+1, preStart+leftSize, inStart, rootIndex-1),
		Right: build(preorder, inorder, preStart+leftSize+1, preEnd, rootIndex+1, inEnd),
	}
}
