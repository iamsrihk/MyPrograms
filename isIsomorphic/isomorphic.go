package main

import "fmt"

func isIsomorphic(str1 string, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}
	s := map[byte]int{}
	t := map[byte]int{}
	for i := range str1 {
		if s[str1[i]] != t[str2[i]] {
			return false
		} else {
			s[str1[i]] = i + 1
			t[str2[i]] = i + 1
		}
	}
	return true
}
func main() {
	fmt.Println("Is isomorphic : ", isIsomorphic("paper", "title"))
}
