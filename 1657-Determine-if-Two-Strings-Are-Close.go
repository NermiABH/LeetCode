func closeStrings(word1 string, word2 string) bool {
    if len(word1) != len(word2) {
        return false
    }
    c1 := make(map[byte]int, len(word1))
    for i := 0; i < len(word1); i++ {
        c1[word1[i]]++ 
    }
    c2 := make(map[byte]int, len(word2))
    for i := 0; i < len(word2); i++ {
        c2[word2[i]]++ 
    }
    for k := range c2 {
        if _, ok := c1[k]; !ok {
            return false
        }
    }
    for k := range c1 {
        if _, ok := c2[k]; !ok {
            return false
        }
    }
    c1Exist := make(map[int]int, len(c1))
    for _, v := range c1{
        c1Exist[v]++
    }
    for _, v := range c2{
        if k := c1Exist[v]; k == 0 {
            return false
        }else {
            c1Exist[v]--
        }
    }
    return true
}