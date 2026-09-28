package task0217

func containsDuplicate(nums []int) bool {
	hash := make(map[int]struct{})
	for _, v := range nums {
		_, ok := hash[v]
		if ok {
			return true
		}
		hash[v] = struct{}{}
	}

	return false
}
