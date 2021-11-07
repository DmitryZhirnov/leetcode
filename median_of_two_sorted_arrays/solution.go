package median_of_two_sorted_arrays

// add Добавление числа в отсортированный массив
func add(nums []int, item int) []int {

	for index, num := range nums {
		if num < item {
			continue
		}
		nums = append(nums[:index+1], nums[index:]...)
		nums[index] = item
		return nums
	}
	return append(nums, item)
}

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	for _, item := range nums2 {
		nums1 = add(nums1, item)
	}

	arraysLen := len(nums1)
	if arraysLen%2 == 0 {
		index := arraysLen / 2
		return float64(nums1[index-1]+nums1[index]) / 2
	}
	return float64(nums1[arraysLen/2])
}
