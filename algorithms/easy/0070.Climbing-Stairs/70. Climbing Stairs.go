package main

func climbStairs(n int) int {
	zero := 1
	one := 1
	next := 0

	for i := 1; i < n; i++ {
		next = zero + one
		zero = one
		one = next
	}

	return one
}

func main() {
	println(climbStairs(2))
	println(climbStairs(3))
}
