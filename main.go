package main

import "fmt"

func main() {
	result := anagram("man", "nak")
	fmt.Println("is angram", result)
}

func anagram(str1, str2 string) bool {
	map1 := make(map[byte]int)
	map2 := make(map[byte]int)
	if len(str1) != len(str2) {
		return false
	}
	for i := 0; i < len(str1); i++ {
		map1[str1[i]]++
		map2[str2[i]]++
	}
	for val := range map1 {
		if map1[val] != map2[val] {
			return false
		}
	}
	return true
}
func isangram(str1, str2 string) bool {
	if len(str1) != len(str2) {
		return false
	}

	for i := 0; i < len(str1); i++ {
	}
}
