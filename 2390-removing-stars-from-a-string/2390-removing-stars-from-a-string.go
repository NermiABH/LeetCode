import "strings"; 

func removeStars(s string) string {
    r := make([]rune, 0)

    for _, v := range s {
        if v == '*'{
            r = r[:len(r)-1]   
        }else {
            r = append(r, v)
        }
    }

    return string(r)
}