package task0202

func isHappy(n int) bool {
	slow := n
	fast := getNext(n)
	for fast != 1 && slow != fast {
		slow = getNext(slow)
		fast = getNext(getNext(fast))
	}
	return fast == 1
}

func getNext(n int) int {
	var result int
	for n > 0 {
		d := n % 10
		result = result + d*d
		n /= 10
	}
	return result
}
