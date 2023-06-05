package leetcode

import (
	"strings"
)

// 151. Reverse Words in a String
func reverseWords(s string) string {
	ans := strings.Fields(s)

	l, r := 0, len(ans)-1
	for l < r {
		temp := string(ans[l])
		ans[l] = string(ans[r])
		ans[r] = temp
		r--
		l++
	}

	return strings.Join(ans[:], " ")
}
