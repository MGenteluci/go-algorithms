package leetcode

// 605. Can Place Flowers
func canPlaceFlowers(flowerbed []int, n int) bool {
	for i := 0; i < len(flowerbed); i++ {
		if flowerbed[i] == 1 {
			continue
		} else if i == 0 {
			if i == len(flowerbed)-1 || flowerbed[i+1] == 0 {
				flowerbed[i] = 1
				n -= 1
			}
		} else if i == len(flowerbed)-1 && flowerbed[i+-1] == 0 {
			flowerbed[i] = 1
			n -= 1
		} else if flowerbed[i-1] == 0 && flowerbed[i+1] == 0 {
			flowerbed[i] = 1
			n -= 1
		}
	}

	return n <= 0
}
