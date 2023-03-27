package reverser

// reverse a string using a for loop and a rune
func Reverse(s string) string {
	var r []rune
	for _, c := range s {
		r = append([]rune{c}, r...)
	}
	return string(r)
}

// func Reverse(s string) string {
// 	r := []rune(s)
// 	for i, j := 0, len(r)-1; i < len(r)/2; i, j = i+1, j-1 {
// 		r[i], r[j] = r[j], r[i]
// 	}
// 	return string(r)
// }
