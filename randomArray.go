package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())

	arr := make([]int, 20000000)
	for i := 0; i < 20000000; i++ {
		arr[i] = rand.Intn(100000) // Generate Random number from 0 - (your input)
	}

	strArr := make([]string, len(arr))
	for i, num := range arr {
		strArr[i] = strconv.Itoa(num)
	}

	result := strings.Join(strArr, " ") // add space

	// write data to file
	err := ioutil.WriteFile("A FILE", []byte(result), 0644) //Change A FILE to the target file without deleting "", Example: "C:/Users/Dafi/Documents/VSCode/RandomNumber.txt"
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	fmt.Println("File has been written.")
}
