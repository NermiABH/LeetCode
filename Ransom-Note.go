func canConstruct(ransomNote string, magazine string) bool {
	var note [26]int
    length := len(ransomNote)
    for i := range ransomNote {
        note[int(byte(ransomNote[i])-'a')]++
    }
    for i := range magazine {
        v := note[int(byte(magazine[i])-'a')]
        if v != 0 {
            note[int(byte(magazine[i])-'a')]--
            length--
            if length == 0 {
                return true
            }
        }
    }
	return false
}
