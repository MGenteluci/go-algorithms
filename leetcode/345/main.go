package leetcode

import "strings"

func isVowel(s string) bool {
	return s == "a" || s == "e" || s == "i" || s == "o" || s == "u" || s == "A" || s == "E" || s == "I" || s == "O" || s == "U"
}

// 345. Reverse Vowels of a String
// Time complexity: O(n)
func reverseVowels(s string) string {
	ans := strings.Split(s, "")

	isVowelLeft, isVowelRight := false, false
	l, r := 0, len(s)-1
	for l < r {
		leftLetter := string(s[l])
		rightLetter := string(s[r])

		if isVowel(leftLetter) {
			isVowelLeft = true
		} else {
			isVowelLeft = false
			ans[l] = leftLetter
			l++
		}

		if isVowel(rightLetter) {
			isVowelRight = true
		} else {
			isVowelRight = false
			ans[r] = rightLetter
			r--
		}

		if isVowelLeft && isVowelRight {
			ans[l] = rightLetter
			ans[r] = leftLetter
			l++
			r--
		}

	}

	return strings.Join(ans[:], "")
}
