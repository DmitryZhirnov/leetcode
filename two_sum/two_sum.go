package two_sum

func TwoSum(nums []int, target int) []int {
	var result []int

	for i := 0; len(nums) > i; i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = []int{i, j}
				return result
			}
		}
	}
	return result
}

func TwoSumWithHashMap(nums []int, target int) []int {
	var result []int
	hashMap := make(map[int]int)

	for i := 0; len(nums) > i; i++ {
		if _, ok:= hashMap[nums[i]]; ok {
			return []int{hashMap[nums[i]], i}
		}
		hashMap[target - nums[i]] = i
	}
	return result
}
