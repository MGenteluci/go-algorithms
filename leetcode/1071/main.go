package leetcode

// 1071. Greatest Common Divisor of Strings
func gcdOfStrings(str1 string, str2 string) string {
	if str1+str2 == str2+str1 {
		return str1[:gcd(len(str1), len(str2))]
	}

	return ""
}

func gcd(str1Len int, str2Len int) int {
	lowestLen := str1Len
	if str2Len < lowestLen {
		lowestLen = str2Len
	}

	for i := lowestLen; i > 0; i-- {
		if str1Len%i == 0 && str2Len%i == 0 {
			return i
		}
	}

	return 0
}
