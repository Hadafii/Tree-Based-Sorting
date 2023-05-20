package main

import (
	"fmt"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Insert(root *Node, value int) *Node {
	if root == nil {
		return &Node{Value: value, Left: nil, Right: nil}
	}

	if value < root.Value {
		root.Left = Insert(root.Left, value)
	} else {
		root.Right = Insert(root.Right, value)
	}

	return root
}

func InorderTraversal(root *Node, result []int) []int {
	if root != nil {
		result = InorderTraversal(root.Left, result)
		result = append(result, root.Value)
		result = InorderTraversal(root.Right, result)
	}
	return result
}

func TreeSort(arr []int) []int {
	var root *Node

	for _, value := range arr {
		root = Insert(root, value)
	}

	result := make([]int, 0)
	result = InorderTraversal(root, result)

	return result
}

func main() {
	arr := []int{}

	sortedArr := TreeSort(arr)

	fmt.Println("Sorted Array:", sortedArr)
}
