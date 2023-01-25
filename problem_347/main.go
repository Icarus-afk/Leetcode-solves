package main

import (
	"fmt"
	"sort"
)

func main(){
	var n int
	var x int 
	var k int 
	var arr []int
	var result []int
	fmt.Println("Enter size of array: ")
	fmt.Scan(&n)
	fmt.Println("Enter elements: ")
	for i:=0;i<n;i++{
		fmt.Scan(&x)
		arr = append(arr, x)
	}
	fmt.Println(arr)
	fmt.Println("Enter the value of K: ")
	fmt.Scan(&k)

	result = topKFrequent(arr, k)
	fmt.Println(result)

}
func topKFrequent(nums []int, k int)[]int{
	answer := []int{}

	var mmp = make(map[int]int)

	for i:=0;i<len(nums); i++{
		mmp[nums[i]] = mmp[nums[i]]+1
	}
	keys:= make([]int, 0,len(mmp))

	for key:= range mmp{
		keys = append(keys, key)
	}
    sort.SliceStable(keys, func(i, j int) bool{
        return mmp[keys[i]] > mmp[keys[j]]
    })

	for j:= 0; j<k; j++{
		answer = append(answer, keys[j])
	}
	return answer
}