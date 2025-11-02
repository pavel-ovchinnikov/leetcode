package divisibleandnondivisiblesumsdifference

func differenceOfSums(n int, m int) int {
	sum1, sum2 := 0, 0

	for i := 0; i <= n; i++ {
		if i%m == 0 {
			sum2 += i
		} else {
			sum1 += i
		}
	}

	return sum1 - sum2
}
