package main

import "fmt"

func main() {
	// a, b := "11", "1"
	a, b := "1010", "1011"
	// a, b := "1001", "11111"
	fmt.Printf("%s\n", addBinary(a, b))
}

func addBinary(a string, b string) string {
	var v1, v2 []byte
	if len(a) > len(b) {
		v1, v2 = invert([]byte(a)), invert([]byte(b))
	} else {
		v1, v2 = invert([]byte(b)), invert([]byte(a))
	}

	res := make([]byte, len(v1)+1)

	var carry bool
	var x1, x2, x byte
	for i := 0; i < len(v2); i++ {
		x1, x2 = v1[i], v2[i]
		switch {
		case x1 == '0' && x2 == '0':
			if carry {
				x = '1'
				carry = false
			} else {
				x = '0'
			}
		case x1 == '1' && x2 == '1':
			if carry {
				x = '1'
			} else {
				x = '0'
			}
			carry = true
		default:
			if carry {
				x = '0'
			} else {
				x = '1'
			}
		}
		res[i] = x
	}

	for i := len(v2); i < len(v1); i++ {
		x1 = v1[i]
		if carry {
			if x1 == '0' {
				x = '1'
				carry = false
			} else {
				x = '0'
			}
		} else {
			x = x1
		}
		res[i] = x
	}

	if carry {
		res = append(res, '1')
	}

	return string(invert(res))
}

func invert(s []byte) []byte {
	r := make([]byte, 0, len(s))
	for i := len(s) - 1; i >= 0; i-- {
		r = append(r, s[i])
	}
	return r
}
