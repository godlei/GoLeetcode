package leetcode

import (
	"fmt"
	"sort"
)

func Problem1() {
	nums := []int{-1, -2, -3, -4, -5}
	target := -8

	res := twoSum(nums, target)
	fmt.Println(res)
}
func twoSum(nums []int, target int) []int {
	numsC := copy(nums)
	sort.Ints(numsC)
	length := len(numsC)
	start := 0
	end := length - 1
	//fmt.Println(numsC)
	for {
		if start >= end {
			break
		}
		if numsC[start]+numsC[end] == target {
			n1 := numsC[start]
			n2 := numsC[end]

			var from, to int
			for i := 0; i < len(nums); i++ {
				if n1 == nums[i] {
					from = i
					break
				}
			}
			for i := 0; i < len(nums); i++ {
				if n2 == nums[i] {
					if from != i {
						to = i
						break
					}
				}
			}
			if from > to {
				from, to = to, from
			}
			return []int{from, to}
		} else if numsC[start]+numsC[end] < target {
			start++
		} else {
			end--
		}
	}
	return nil
}

func copy(nums []int) (res []int) {
	for _, n := range nums {
		res = append(res, n)
	}
	return
}
