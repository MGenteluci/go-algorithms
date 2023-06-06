package leetcode

// 238. Product of Array Except Self
// Time complexity O(n)
func productExceptSelf(nums []int) []int {
	N := len(nums)
	ans := make([]int, N)

	ans[0] = 1
	for i := 1; i < N; i++ {
		ans[i] = ans[i-1] * nums[i-1]
	}

	suff := 1
	for i := N - 2; i >= 0; i-- {
		suff = suff * nums[i+1]
		ans[i] = ans[i] * suff
	}

	return ans
}
