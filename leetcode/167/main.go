package main

// 167. Two Sum II - Input Array Is Sorted
// Time complexity: O(n)
func twoSum(numbers []int, target int) []int {
	ans := []int{}
	l, r := 0, len(numbers)-1

	for {
		if numbers[l]+numbers[r] == target {
			ans = append(ans, l+1, r+1)
			break
		} else if numbers[l]+numbers[r] < target {
			l++
		} else {
			r--
		}
	}

	return ans
}

func main() {}
