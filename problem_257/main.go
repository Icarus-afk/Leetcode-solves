package main

import "fmt"

func main() {
	var i, n, x int
	var arr []int
	fmt.Println("Enter size of array: ")
	fmt.Scan(&n)
	fmt.Println("Enter a array: ")
	for i = 0; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(arr)
	fmt.Println(containsDuplicate(arr))
}

func containsDuplicate(nums []int) bool {
	duplicates := make(map[int]bool)
	for _, num := range nums {
		if duplicates[num] == true {
			return true
		} else {
			duplicates[num] = true
		}
	}
	return false
}
