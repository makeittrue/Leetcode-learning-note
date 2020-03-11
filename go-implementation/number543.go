package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if (root.Left == nil) && (root.Right == nil){
		return 0
	}else{
		max := 0
		setDepth(root, &max)
		return max
	}
}

func setDepth(root *TreeNode, max *int) int{
	if root != nil{
		right := setDepth(root.Right, max)
		left := setDepth(root.Left, max)
		if right + left > *max {
			*max = right + left
		}
		if left > right {
			return left +1
		}else{
			return right +1
		}
	}
	return 0
}