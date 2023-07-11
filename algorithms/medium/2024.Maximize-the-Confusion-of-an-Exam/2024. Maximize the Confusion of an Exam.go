package leetcode

func maxConsecutiveAnswers(answerKey string, k int) int {
	if len(answerKey) < 2 {
		return len(answerKey)
	}

	size := len(answerKey)
	maxRepeatCharCount := 0
	frequencyMap := make(map[byte]int)
	left, right := 0, 0
	res := 0

	for ; right < size; right++ {
		rightChar := answerKey[right]
		frequencyMap[rightChar] += 1

		if frequencyMap[rightChar] > maxRepeatCharCount {
			maxRepeatCharCount = frequencyMap[rightChar]
		}

		for right-left+1-maxRepeatCharCount > k {
			leftChar := answerKey[left]
			frequencyMap[leftChar]--
			left++
		}

		if right-left+1 > res {
			res = right - left + 1
		}
	}
	return res
}
