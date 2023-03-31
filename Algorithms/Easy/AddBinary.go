package main

import (
	"strconv"
)

func addBinary(a string, b string) string {
	var s string
	i, j := len(a)-1, len(b)-1
	if i < j {
		i, j = j, i
		a, b = b, a
	}
	add := 0
	for j >= 0 {
		var elem int
		if a[i] == '1' && b[j] == '1' {
			if add == 1 {
				elem = add
			} else {
				add++
			}
		} else if a[i] == '0' && b[j] == '0' {
			elem, add = add, 0
		} else if add != 1 {
			elem = 1
		}
		s = strconv.Itoa(elem) + s
		i--
		j--
	}
	for i >= 0 && add == 1 {
		var elem int
		if a[i] == '0' {
			i--
			elem = 1
			add = 0
			s = strconv.Itoa(elem) + s
			break
		}
		s = strconv.Itoa(elem) + s
		i--
	}
	if i >= 0 {
		s = a[:i+1] + s
	}
	if add == 1 {
		s = "1" + s
	}
	return s
}
