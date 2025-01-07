func maxVowels(s string, k int) int {
    m := 0
    for i := range s[:k] {
        if isVowel(s[i]) {
           m++
        }
    }
    l := m
    for i := k; i < len(s); i++{
        if isVowel(s[i-k]){
            l--
        }
        if isVowel(s[i]){
            l++
            if m <l{
                m = l
            }
        }
    }
    return m
}

func isVowel(c byte) bool {
    if c < 'a' { 
        c += 32 
    }
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u'
}