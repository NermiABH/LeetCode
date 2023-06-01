package main

import "fmt"

func isAnagram(s string, t string) bool {
	if len(s) != len(t) {
		return false
	}
	symbols := make(map[rune]uint, len(s))
	for _, v := range s {
		symbols[v]++
	}
	for _, v := range t {
		if value, ok := symbols[v]; ok || value == 0 {
			return false
		}
		symbols[v]--
	}
	return true
}

//func isAnagram(s string, t string) bool {
//	runeS, runeT := []rune(s), []rune(t)
//	sort.Slice(runeS, func(i, j int) bool {
//		return runeS[i] < runeS[j]
//	})
//	sort.Slice(runeT, func(i, j int) bool {
//		return runeT[i] < runeT[j]
//	})
//	for i, v := range runeS {
//		if v != runeT[i] {
//			return false
//		}
//	}
//	return true
//}

func main() {
	fmt.Println(string(rune(-23123)))
	fmt.Println(isAnagram("пятак", "пятка"))
}
