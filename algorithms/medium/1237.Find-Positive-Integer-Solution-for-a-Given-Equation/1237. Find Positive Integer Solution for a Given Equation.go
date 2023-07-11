package main

/**
 * This is the declaration of customFunction API.
 * @param  x    int
 * @param  x    int
 * @return 	    Returns f(x, y) for any given positive integers x and y.
 *			    Note that f(x, y) is increasing with respect to both x and y.
 *              i.e. f(x, y) < f(x + 1, y), f(x, y) < f(x, y + 1)
 */

func findSolution(customFunction func(int, int) int, z int) [][]int {
	res := [][]int{}

	for x := 1; x <= z; x++ {
		b, e := 1, z
		for b <= e {
			m := b + (e-b)/2
			if customFunction(x, m) == z {
				res = append(res, []int{x, m})
				e = m - 1
			} else if customFunction(x, m) > z {
				e = m - 1
			} else {
				b = m + 1
			}
		}
	}
	return res
}

func main() {
	findSolution(func(a, b int) int { return a + b }, 10)
}
