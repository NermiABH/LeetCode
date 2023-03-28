package main

func canConstruct(ransomNote string, magazine string) bool {
	magazineChar := mapOfSlice(magazine)
	for _, val := range ransomNote {
		if count, exist := magazineChar[string(val)]; count == 0 || !exist {
			return false
		}
		magazineChar[string(val)] -= 1
	}
	return true
}

func mapOfSlice(s string) map[string]int {
	charCount := make(map[string]int)
	for _, val := range s {
		charCount[string(val)] += 1
	}
	return charCount
}
