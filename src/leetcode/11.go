package leetcode

import "fmt"

func Problem11() {
	array := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	fmt.Println(maxArea(array))
}

func maxArea(height []int) int {
	if len(height) < 2 {
		return 0
	}

	if len(height) == 2 {
		min := height[0]
		if height[1] < min {
			min = height[1]
		}
		return min
	}

	from := 0
	to := len(height) - 1
	maxArea := 0

	for from < to {
		min := height[from]
		if height[to] < min {
			min = height[to]
		}
		area := min * (to - from)

		if area > maxArea {
			maxArea = area
		}

		if height[from] < height[to] {
			from++
		} else {
			to--
		}
	}

	return maxArea
}
