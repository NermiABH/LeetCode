func groupAnagrams(strs []string) [][]string {
    set := make(map[[26]byte][]string)

    for _, v := range strs {
        var chars [26]byte
        for i := 0; i < len(v); i++ {
            chars[v[i]-'a']++
        }
        set[chars] = append(set[chars], v)
    }

    resp := make([][]string, 0, len(set))
    for _, v := range set{
        resp = append(resp, v)
    }

    return resp
}