package task0001

func twoSum(nums []int, target int) []int {
	hash := make(map[int]int, 0)
	for i, num := range nums {
		if ii, ok := hash[target-num]; ok {
			return []int{ii, i}
		}
		hash[num] = i
	}
	return nil
}
