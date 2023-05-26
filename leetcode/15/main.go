package main

import (
	"sort"
)

// 15. 3Sum
func threeSum(nums []int) [][]int {
	sort.Ints(nums)
	ans := [][]int{}

	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		l, r := i+1, len(nums)-1
		if l == r {
			break
		}

		for l < r {
			threeSum := nums[i] + nums[l] + nums[r]
			if threeSum > 0 {
				r--
			} else if threeSum < 0 {
				l++
			} else {
				ans = append(ans, []int{nums[i], nums[l], nums[r]})
				l++

				for nums[l] == nums[l-1] && l < r {
					l++
				}
			}
		}
	}

	return ans
}

func main() {}
