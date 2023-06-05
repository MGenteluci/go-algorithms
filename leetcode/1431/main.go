package leetcode

// 1431. Kids With the Greatest Number of Candies
// Time complexity O(n)
func kidsWithCandies(candies []int, extraCandies int) []bool {
	ans := make([]bool, len(candies))
	highest := 0

	for i := 0; i < len(candies); i++ {
		if candies[i] > highest {
			highest = candies[i]
		}
	}

	for i := 0; i < len(candies); i++ {
		if candies[i]+extraCandies >= highest {
			ans[i] = true
		}
	}

	return ans
}
