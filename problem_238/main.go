package main

import (
	"fmt"
)

func main(){
	var n int
	var x int
	var arr []int
	fmt.Println("Enter the size of array: ")
	fmt.Scan(&n)
	fmt.Println("Enter array elements: ")
	for i:=0; i<n; i++{
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(arr)
	fmt.Println(productExceptSelf(arr))
}

func productExceptSelf(nums []int) []int{
    res := make([]int, len(nums))

    prefix := 1
    for i := 0; i < len(res); i++ {
        res[i] = prefix
        prefix *= nums[i]
    }

    postfix := 1
    for i := len(res) - 1; i >= 0; i-- {
        res[i] *= postfix
        postfix *= nums[i]
    }

    return res
}