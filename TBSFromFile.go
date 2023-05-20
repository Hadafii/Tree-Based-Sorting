package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
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
	// Read file
	start := time.Now()
	content, err := ioutil.ReadFile("Read File") //Change Read File to the Reading target file without deleting "", Example: "C:/Users/Dafi/Documents/VSCode/RandomNumber.txt"
	if err != nil {
		log.Fatal(err)
	}

	// changing file text to integer
	scanner := bufio.NewScanner(strings.NewReader(string(content)))
	scanner.Split(bufio.ScanWords)

	arr := []int{}
	for scanner.Scan() {
		num, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		arr = append(arr, num)
	}

	sortedArr := TreeSort(arr)

	// Changing array from sorted number to string with space
	strArr := make([]string, len(sortedArr))
	for i, num := range sortedArr {
		strArr[i] = strconv.Itoa(num)
	}
	result := strings.Join(strArr, " ")

	// Write data to a file
	err = ioutil.WriteFile("Extracted File", []byte(result), 0644) //Change Extracted File to the target file without deleting "", Example: "C:/Users/Dafi/Documents/VSCode/Sorted.txt"
	if err != nil {
		log.Fatal(err)
	}
	elapsed := time.Since(start)

	fmt.Printf("Sorted Array has been written to file. Time elapsed: %s", elapsed)
}
