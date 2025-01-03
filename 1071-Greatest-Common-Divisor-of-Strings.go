func gcdNum(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func gcdOfStrings(str1 string, str2 string) string {
	if len(str1) < len(str2) {
		str1, str2 = str2, str1
	}
Main:
	for i := gcdNum(len(str2), len(str1)); i > 0; i /= 2 {
		gcd := str2[:i]
		for j := 0; ; j += i {
			if gcd != str2[j:j+i] {
				continue Main
			}
			if j+i >= len(str2) {
				break
			}
		}
		for j := 0; ; j += i {
			if gcd != str1[j:i+j] {
				continue Main
			}
			if j+i >= len(str1) {
				break
			}
		}
		return gcd
	}
	return ""
}