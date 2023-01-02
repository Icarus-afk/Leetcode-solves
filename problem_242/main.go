package main

import "fmt"

func main() {
	var s, t string
	fmt.Println("Enter a string: ")
	fmt.Scanln(&s)
	fmt.Println("Enter another string to check anagram: ")
	fmt.Scanln(&t)
	fmt.Println(s, t)
	fmt.Println(isAnagram(s, t))
}
func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	rec := [26]int{}
	for _, v := range s {
		rec[v-'a']++
	}
	for _, v := range t {
		rec[v-'a']--
	}
	for _, v := range rec {
		if v != 0 {
			return false
		}
	}
	return true
}
