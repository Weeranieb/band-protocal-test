package rescuechicken

// return max member that contains
// two pointers track
func RescueChicken(chickenNo, length int, position []int) int {
	// edge case
	if chickenNo == 0 {
		return 0
	}

	// min position is on the left of the array
	left := 0
	maxCount := 0

	// right pointer
	for right := 0; right < chickenNo; right++ {
		// to move left pointer to the right if the distance is too long
		for position[right]-position[left] > length-1 {
			left++
		}

		// always update maxCount
		maxCount = max(maxCount, right-left+1)
	}

	return maxCount
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
