package main

import "fmt"

func main() {
	var n, x, target int
	var arr []int
	fmt.Println("Enter Array size: ")
	fmt.Scan(&n)
	fmt.Println("Enter the array: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(arr)
	fmt.Println("Enter target value: ")
	fmt.Scan(&target)
	fmt.Println(twoSum(arr, target))

}

func twoSum(nums []int, target int) []int {
	var returnarr []int
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				returnarr = append(returnarr, i, j)
			}
		}
	}
	return returnarr
}
