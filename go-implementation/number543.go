package main

type TreeNode struct {
     Val int
     Left *TreeNode
     Right *TreeNode
}
var max = 0

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil{
		return 0
	}
	if (root.Left == nil) && (root.Right == nil){
		return 0
	}else{
		setDepth(root)
		return max
	}
}

func setDepth(root *TreeNode) int{
	if root != nil{
		right := setDepth(root.Right)
		left := setDepth(root.Left)
		if right + left > max {
			max = right + left
		}
		if left > right {
			return left +1
		}else{
			return right +1
		}
	}
	return 0
}