package main

import "fmt"

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func inOrder(root *Node) {
	if root == nil {
		return
	}

	inOrder(root.Left)
	fmt.Print(root.Val, " ")
	inOrder(root.Right)
}

func insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Val: value, Left: nil, Right: nil}
	}

	if value < root.Val {
		root.Left = insert(root.Left, value)
	} else {
		root.Right = insert(root.Right, value)
	}

	return root
}

func insertArray(root *Node, arr []int) *Node {
	for _, num := range arr {
		root = insert(root, num)
	}

	return root
}

func countNodes(tree *Node) int {
	if tree == nil {
		return 0
	}

	return countNodes(tree.Left) + countNodes(tree.Right) + 1
}

func getElements(tree *Node, aux []int, index *int) {
	if tree != nil {
		getElements(tree.Left, aux, index)
		aux[*index] += tree.Val
		*index++
		getElements(tree.Right, aux, index)
	}
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var tree *Node

	tree = insertArray(tree, nums1)
	tree = insertArray(tree, nums2)

	if tree != nil {
		size := countNodes(tree)
		aux := make([]int, size, size)

		index := 0

		getElements(tree, aux, &index)

		if size%2 != 0 {
			return float64(aux[size/2])
		} else {
			return float64(aux[(size-1)/2]+aux[size/2]) / 2.0
		}

	}

	return 0
}

func main() {
	fmt.Println(findMedianSortedArrays([]int{1, 2}, []int{3, 4}))
}
