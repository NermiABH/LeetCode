func reverseVowels(s string) string {
	b := []byte(s)
	for i, j := 0, len(b)-1; i < j; {
		if !isVowel(s[i]) {
			i++
			continue
		}
		if !isVowel(s[j]) {
			j--
			continue
		}
		b[i], b[j] = b[j], b[i]
		i++
		j--
	}
	return *(*string)(unsafe.Pointer(&b))
}

func isVowel(c byte) bool {
    if c < 'a' { 
        c += 32 
    }
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}