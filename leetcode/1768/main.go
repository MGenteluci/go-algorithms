package main

// 1768. Merge Strings Alternately
// Time complexity O(n)
func mergeAlternately(word1 string, word2 string) string {
	ans, i, j, word1Len, word2Len := "", 0, 0, len(word1), len(word2)
	
	for i < word1Len && j < word2Len {
		ans = ans + string(word1[i])
		ans = ans + string(word2[j])
		i++
		j++
	}

	for i < word1Len {
		ans = ans + string(word1[i])
		i++
	}

	for j < word2Len {
		ans = ans + string(word2[j])
		j++
	}

	return ans
}

func main() {}
