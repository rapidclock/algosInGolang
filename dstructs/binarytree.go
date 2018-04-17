package dstructs

import (
	"strings"
	"strconv"
	"../genutils"
)

type TreeNode struct {
	Left  *TreeNode
	Val   int
	Right *TreeNode
}

func NewTreeNode(val int) *TreeNode {
	node := new(TreeNode)
	node.Val = val
	return node
}

func (root *TreeNode) String() string {
	return root.SerializeTree()
}

func (root *TreeNode) MinDepth() int {
	if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return 1 + root.Right.MinDepth()
	} else if (root.Right == nil) {
		return 1 + root.Left.MinDepth()
	} else {
		return 1 + genutils.MinOfTwo(root.Left.MinDepth(), root.Right.MinDepth())
	}
}

func (root *TreeNode) MaxDepth() int {
	if root == nil {
		return 0
	} else {
		return 1 + genutils.MaxOfTwo(root.Left.MaxDepth(), root.Right.MaxDepth())
	}
}

func (root *TreeNode) CountNodes() int {
	var leftCnt, rightCnt int
	if root.Left != nil {
		leftCnt = root.Left.CountNodes()
	}
	if root.Right != nil {
		rightCnt = root.Right.CountNodes()
	}
	return 1 + leftCnt + rightCnt
}

func IsBalancedBinaryTree(root *TreeNode) bool {
	ok, _ := checkBalance(root)
	return ok
}

func checkBalance(root *TreeNode) (bool, int) {
	if root == nil {
		return true, 0
	} else {
		leftCheck, leftDepth := checkBalance(root.Left)
		rightCheck, rightDepth := checkBalance(root.Right)
		balanceCheck := leftCheck && rightCheck && (genutils.Abs(leftDepth-rightDepth) <= 1)
		maxDepth := 1 + genutils.MaxOfTwo(leftDepth, rightDepth)
		return balanceCheck, maxDepth
	}
}

func (root *TreeNode) PreOrder() []int {
	soln := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		if node != nil {
			soln = append(soln, node.Val)
			if node.Right != nil {
				stack = append(stack, node.Right)
			}
			node = node.Left
		} else {
			node, stack = stack[len(stack)-1], stack[:len(stack)-1]
		}
	}
	return soln
}

func (root *TreeNode) InOrder() []int {
	soln := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			node, stack = stack[len(stack)-1], stack[:len(stack)-1]
			soln = append(soln, node.Val)
			node = node.Right
		}
	}
	return soln
}

func (root *TreeNode) PostOrder() []int {
	soln := make([]int, 0)
	stack := make([]*TreeNode, 0)
	node := root
	for len(stack) > 0 || node != nil {
		if node != nil {
			stack = append(stack, node)
			node = node.Left
		} else {
			temp := stack[len(stack)-1].Right
			if temp != nil {
				node = temp
			} else {
				temp, stack = stack[len(stack)-1], stack[:len(stack)-1]
				soln = append(soln, temp.Val)
				for len(stack) > 0 && temp == stack[len(stack)-1].Right {
					temp, stack = stack[len(stack)-1], stack[:len(stack)-1]
					soln = append(soln, temp.Val)
				}
			}
		}
	}
	return soln
}

func (root *TreeNode) SerializeTree() string {
	soln := make([]string, 0)
	codify(root, &soln)
	return strings.Join(soln, ",")
}

func codify(root *TreeNode, soln *[]string) {
	if root == nil {
		*soln = append(*soln, "#")
	} else {
		*soln = append(*soln, strconv.Itoa(root.Val))
		codify(root.Left, soln)
		codify(root.Right, soln)
	}
}

func BuildTree(serial string) *TreeNode {
	nodes := strings.Split(serial, ",")
	root := decodify(&nodes)
	return root
}

func decodify(nodes *[]string) *TreeNode {
	if len(*nodes) == 0 {
		return nil
	}
	var nodeVal string
	nodeVal, *nodes = (*nodes)[0], (*nodes)[1:]
	if nodeVal == "#" {
		return nil
	} else {
		nodeIntVal, _ := strconv.Atoi(nodeVal)
		node := NewTreeNode(nodeIntVal)
		node.Left = decodify(nodes)
		node.Right = decodify(nodes)
		return node
	}
}
