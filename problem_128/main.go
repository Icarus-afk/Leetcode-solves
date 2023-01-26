package main

import (
	"fmt"
)

func main(){
	var n int
	var arr []int
	var x int
	fmt.Print("Enter the size of array: ")
	fmt.Scan(&n)
	fmt.Print("Enter array elements: ")
	for i:=0; i<n;i++{
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(arr)
	fmt.Println(longestConsecutive(arr))
}

func longestConsecutive(nums []int)int{

	mp := make(map[int]int)

	for _,v:=range nums{
		mp[v] = 1
	}
	
	ans:=0
	for _, v:=range nums{
		if mp[v-1]==1{
			continue
		}
		y:=v+1
		for mp[y]==1{
			y++
		}
		if y-v > ans{
			ans = y-v
		}
	}
	return ans
}