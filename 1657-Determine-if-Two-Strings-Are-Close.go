func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	chars1, chars2 := [26]int{}, [26]int{}
	for i := range word1 {
		chars1[word1[i]-'a']++
        chars2[word2[i]-'a']++
	}
    n := make(map[int]int, 26)
	for i := range chars1{
        if (chars1[i] > 0 && chars2[i] == 0) || (chars1[i] == 0 && chars2[i] > 0){
            return false
        }else if chars1[i] > 0 && chars2[i] > 0 {
            n[chars1[i]]++
            n[chars2[i]]--
        }
    }
    for _, v := range n {
        if v != 0 {
            return false
        }
    }
	return true
}
