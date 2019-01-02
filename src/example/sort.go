package example

import "fmt"

func SortExample() {
	nums := []int{13, 65, 97, 76, 38, 27, 49}

	//nums = []int{5,4,3,2,1}
	fmt.Println("排序前:")
	fmt.Println(nums)
	selectSort(nums)
	insertSort(nums)
	bubleSort(nums)
	mergeSort(nums)
	quickSort(nums)
}

//O(n*n) O(n*n) O(n*n)
func selectSort(nums []int) {
	fmt.Println("选择排序:")
	if nums == nil || len(nums) <= 1 {
		fmt.Println(nums)
		return
	}

	for i := 0; i < len(nums)-1; i++ {
		minPos := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minPos] {
				minPos = j
			}
		}
		if minPos > i {
			nums[i], nums[minPos] = nums[minPos], nums[i]
		}
	}
	fmt.Println(nums)
}

//O(n) O(n*n) O(n*n)
func insertSort(nums []int) {
	fmt.Println("插入排序:")
	if nums == nil || len(nums) <= 1 {
		fmt.Println(nums)
		return
	}

	for i := 1; i < len(nums); i++ {
		n := nums[i]
		found := false
		for j := i - 1; j >= 0; j-- {
			if nums[j] > n {
				nums[j+1] = nums[j]
			} else {
				nums[j+1] = n
				found = true
				break
			}
		}
		if !found {
			nums[0] = n
		}
	}
	fmt.Println(nums)
}

//O(n) O(n*n) O(n*n)
func bubleSort(nums []int) {
	fmt.Println("冒泡排序:")
	if nums == nil || len(nums) <= 1 {
		fmt.Println(nums)
		return
	}

	for j := len(nums) - 1; j > 0; j-- {
		swap := false
		for i := 0; i < j; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				swap = true
			}
		}
		if !swap {
			break
		}
	}

	fmt.Println(nums)
}

//O(n*logn) O(n*logn) O(n*logn)
func mergeSort(nums []int) {
	fmt.Println("归并排序:")
	if nums == nil || len(nums) <= 1 {
		fmt.Println(nums)
		return
	}
	//nums = []int{3,48,25}
	//merge(nums,0,1,2)
	doMergeSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}
func merge(nums []int, from int, mid int, to int) {
	var a1, a2 []int

	for i := from; i <= mid; i++ {
		a1 = append(a1, nums[i])
	}

	for i := mid + 1; i <= to; i++ {
		a2 = append(a2, nums[i])
	}
	index1 := 0
	index2 := 0
	for i := from; i <= to; i++ {
		if index1 < len(a1) && index2 < len(a2) {
			if a1[index1] <= a2[index2] {
				nums[i] = a1[index1]
				index1++
			} else {
				nums[i] = a2[index2]
				index2++
			}
		} else {
			if index1 < len(a1) {
				nums[i] = a1[index1]
				index1++
			}
			if index2 < len(a2) {
				nums[i] = a2[index2]
				index2++
			}
		}

	}
}
func doMergeSort(nums []int, from int, to int) {
	if from >= to {
		return
	}
	pos := (from + to) / 2

	doMergeSort(nums, from, pos)
	doMergeSort(nums, pos+1, to)

	merge(nums, from, pos, to)
}

func quickSort(nums []int) {
	fmt.Println("快速排序:")
	if nums == nil || len(nums) <= 1 {
		fmt.Println(nums)
		return
	}

	doQuickSort(nums, 0, len(nums)-1)
	fmt.Println(nums)
}

func doQuickSort(nums []int, from int, to int) {
	if from >= to {
		return
	}
	start := from
	end := to

	val := nums[from]

	for {
		if from == to {
			break
		}

		for {
			if to > from && nums[to] >= val {
				to--
			} else {
				nums[from] = nums[to]
				break
			}
		}

		for {
			if to > from && nums[from] <= val {
				from++
			} else {
				nums[to] = nums[from]
				break
			}
		}
	}
	nums[from] = val
	doQuickSort(nums, start, from-1)
	doQuickSort(nums, from+1, end)
}
