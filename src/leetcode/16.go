package leetcode

import (
	"math"
	"sort"
)

func threeSumClosest(nums []int, target int) int {
	if len(nums) == 3 {
		res := nums[0] + nums[1] + nums[2]
		return res
	}
	sort.Ints(nums)
	//fmt.Println(nums)
	res := nums[0] + nums[1] + nums[2]
	minDiff := math.Abs(float64(res - target))
	index := 0
	for index < len(nums)-2 {
		i := index + 1
		j := len(nums) - 1

		for i < j {

			resTemp := nums[index] + nums[i] + nums[j]
			diff := math.Abs(float64(resTemp - target))

			//fmt.Printf("index:%d,i:%d j:%d,diff:%f,minDiff:%f,resTemp:%d\n",index,i,j,diff,minDiff,resTemp)
			if diff < minDiff {
				res = resTemp
				minDiff = diff
			}
			if resTemp-target < 0 {
				i++
				for i < j && nums[i] == nums[i-1] {
					i++
				}
			} else if resTemp-target > 0 {
				j--
				for i < j && nums[j] == nums[j+1] {
					j--
				}

			} else {
				return target
			}
		}
		index++
	}
	return res
}
