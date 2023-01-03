package main

import "fmt"

func main() {
	var n int
	var str string
	var strarr []string
	fmt.Println("Enter the size of array: ")
	fmt.Scan(&n)
	fmt.Println("Enter the strings: ")
	for i := 0; i < n; i++ {
		fmt.Scan(&str)
		strarr = append(strarr, str)
	}
	fmt.Println(strarr)
	fmt.Print(groupAnagrams(strarr))

}
func groupAnagrams(strs []string) [][]string {
	mp := map[[26]int][]string{}
	for _, s := range strs {
		k := [26]int{}
		for i := 0; i < len(s); i++ {
			k[s[i]-'a'] += 1
		}
		mp[k] = append(mp[k], s)
	}
	returnarr := [][]string{}
	for _, v := range mp {
		returnarr = append(returnarr, v)
	}
	return returnarr
}
