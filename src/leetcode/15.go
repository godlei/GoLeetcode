package leetcode

import (
	"fmt"
	"sort"
)

func Problem15() {
	nums := []int{-1, 0, 1, 2, -1, -4}
	fmt.Println(threeSum(nums))
}

func threeSum(nums []int) (res [][]int) {
	if len(nums) < 3 {
		return
	}
	sort.Ints(nums)
	if len(nums) == 3 {
		if nums[0]+nums[1]+nums[2] == 0 {
			res = append(res, nums)

		}
		return
	}

	for i := 0; i < len(nums)-2; i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		from := i + 1
		to := len(nums) - 1
		target := 0 - nums[i]

		for from < to {
			if nums[from]+nums[to] < target {
				from++
				for from < to && nums[from] == nums[from-1] {
					from++
				}
				if from >= to {
					break
				}
			}

			if nums[from]+nums[to] > target {
				to--
				for to > from && nums[to] == nums[to+1] {
					to--
				}

				if from >= to {
					break
				}
			}

			if nums[from]+nums[to] == target {
				res = append(res, []int{nums[i], nums[from], nums[to]})

				to--
				from++

				for from < to && nums[from] == nums[from-1] {
					from++
				}
				if from >= to {
					break
				}
				for to > from && nums[to] == nums[to+1] {
					to--
				}
				if from >= to {
					break
				}
			}
		}
	}
	return
}
