package main

import "fmt"

type Node struct {
	data   int
	parent *Node
	left   *Node
	right  *Node
}

type BinaryTree struct {
	root *Node
}

func (tree *BinaryTree) InsertItem(i int) {
	if tree.root == nil {
		tree.root = &Node{data: i}
		return
	}
	currentNode := tree.root
	for {
		if i > currentNode.data {
			if currentNode.right == nil {
				currentNode.right = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.right
		} else {
			if currentNode.left == nil {
				currentNode.left = &Node{data: i, parent: currentNode}
				return
			}
			currentNode = currentNode.left
		}
	}
}

func (tree *BinaryTree) SearchItem(i int) (*Node, bool) {
	if tree.root == nil {
		return nil, false
	}
	currentNode := tree.root
	for currentNode != nil {
		if i == currentNode.data {
			return currentNode, true
		} else if i > currentNode.data {
			currentNode = currentNode.right
		} else if i < currentNode.data {
			currentNode = currentNode.left
		}
	}
	return nil, false
}

func searchMatrix(matrix [][]int, target int) bool {
    lengthX := len(matrix)
    for i := 0; i < lengthX; i++ {
        if target >= matrix[i][0] && target <= matrix[i][len(matrix[i]) - 1] {
            tree := &BinaryTree{}
            for _, value := range matrix[i] {
                tree.InsertItem(value)    
            }
            _, hasFound := tree.SearchItem(target)
            return hasFound
        }
    }
    return false
}

func main() {
	matrix := [][]int{{1,3,5,7},{10,11,16,20},{23,30,34,60}}
	fmt.Println(searchMatrix(matrix, 3))
}