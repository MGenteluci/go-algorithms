package main

// 3. Longest Substring Without Repeating Characters
// Time Complexity: O(n)
func lengthOfLongestSubstring(s string) int {
	longest := 0
	ss := map[string]int{}

	for i, j := 0, 0; i < len(s); i++ {
		letter := string(s[i])
		for {
			if _, exist := ss[letter]; !exist {
				break
			}
			delete(ss, string(s[j]))
			j++
		}

		ss[letter] = i
		currLen := i - j + 1
		if currLen > longest {
			longest = currLen
		}
	}

	return longest
}

func main() {}
