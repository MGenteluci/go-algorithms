package main

// 1. Two Sum
// Hashtable solution
// Time complexity O(n)
func twoSum(nums []int, target int) []int {
    hashtable := map[int]int{}

    for i := 0; i < len(nums); i++ {
        complement := target - nums[i]
		if complementIdx, exist := hashtable[complement]; exist {
			return []int{complementIdx, i}
		}

		hashtable[nums[i]] = i
    }

	return []int{0, 0}
}

func main() {}
