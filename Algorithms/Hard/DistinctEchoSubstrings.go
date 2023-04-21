package main

import "fmt"

func distinctEchoSubstrings(text string) int {
	exo := make(map[string]struct{})
	for i := range text {
		for j := i + 1; j-i <= len(text)-j; j++ {
			if t := text[i:j]; t == text[j:2*j-i] {
				exo[t] = struct{}{}
			}
		}
	}
	return len(exo)
}

func main() {
	text := "abcabcabc"
	fmt.Println(distinctEchoSubstrings(text))
}
