package main

import "fmt"

func checkIfCanBreak(s1 string, s2 string) bool {
	m1, m2 := [26]int{}, [26]int{}
	for i := range s1 {
		m1[s1[i]-'a']++
		m2[s2[i]-'a']++
	}
	less, more := 0, 0
	for i := range m1 {
		if less >= 0 {
			less += m2[i] - m1[i]
		}
		if more >= 0 {
			more += m1[i] - m2[i]
		}
	}
	return less >= 0 || more >= 0
}

func main() {
	fmt.Println(checkIfCanBreak("leetcodee", "interview"))
}
