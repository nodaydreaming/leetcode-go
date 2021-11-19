package main

/**
给定一个二叉树，计算 整个树 的坡度 。

一个树的 节点的坡度 定义即为，该节点左子树的节点之和和右子树节点之和的 差的绝对值 。如果没有左子树的话，左子树的节点之和为 0 ；没有右子树的话也是一样。空结点的坡度是
整个树 的坡度就是其所有节点的坡度之和。

来源：力扣（LeetCode）
链接：https://leetcode-cn.com/problems/binary-tree-tilt
著作权归领扣网络所有。商业转载请联系官方授权，非商业转载请注明出处。
*/

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func findTilt(root *TreeNode) int {
	ans := 0
	var dfs func(node *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}
		sumLeft := dfs(root.Left)
		sumRight := dfs(root.Right)
		ans += abs(sumLeft - sumRight)

		return sumLeft + sumRight + root.Val
	}

	dfs(root)

	return ans
}

func abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}

func main() {
	var root *TreeNode

	findTilt(root)
}
