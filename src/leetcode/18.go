package leetcode

import "sort"

func fourSum(nums []int, target int) (result [][]int) {

	if len(nums) < 4 {
		return
	}
	if len(nums) == 4 {
		if nums[0]+nums[1]+nums[2]+nums[3] == target {
			result = append(result, nums)
			return
		} else {
			return
		}

	}

	sort.Ints(nums)

	var first, second, third, forth int
	for first = 0; first < len(nums)-3; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		for second = first + 1; second < len(nums)-2; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			third = second + 1
			forth = len(nums) - 1
			for third < forth {
				if nums[first]+nums[second]+nums[third]+nums[forth] == target {
					array := []int{nums[first], nums[second], nums[third], nums[forth]}
					result = append(result, array)

					third++
					for third < forth && nums[third-1] == nums[third] {
						third++
					}

					forth--
					for third < forth && nums[forth+1] == nums[forth] {
						forth--
					}
				} else if nums[first]+nums[second]+nums[third]+nums[forth] < target {
					third++
					for third < forth && nums[third-1] == nums[third] {
						third++
					}
				} else {
					forth--
					for third < forth && nums[forth+1] == nums[forth] {
						forth--
					}
				}
			}
		}
	}
	return result
}
