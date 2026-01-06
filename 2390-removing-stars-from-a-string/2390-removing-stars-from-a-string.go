import "strings"; 

func removeStars(s string) string {
    stack:=make([]byte, len(s))
    j:=0
    for i:=0; i<len(s); i++ {
        if s[i]!='*' {
            stack[j]=s[i]
            j++
        }
        if s[i]=='*' {
            j--
        }
    }
    return string(stack[:j])
}