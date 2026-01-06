import "strings"; 

func removeStars(s string) string {
    r := make([]byte, 0)

    for i := range s {
        b := s[i]
        if b == '*'{
            r = r[:len(r)-1]   
        }else {
            r = append(r, b)
        }
    }

    return string(r)
}