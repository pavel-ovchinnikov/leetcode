package main

import "fmt"

func sortColors(nums []int) {
	RED, WHITE, BLUE := 0, 1, 2
	idxRed, idxBlue := 0, len(nums)-1

	for i := 0; i <= idxBlue; {
		switch nums[i] {
		case RED:
			nums[i], nums[idxRed] = nums[idxRed], nums[i]
			idxRed++
			i++
		case WHITE:
			i++
		case BLUE:
			nums[i], nums[idxBlue] = nums[idxBlue], nums[i]
			idxBlue--
		}
	}
}

func main() {
	{
		nums := []int{2, 0, 2, 1, 1, 0}
		sortColors(nums)
		fmt.Println(nums)
	}
	{
		nums := []int{2, 0, 1}
		sortColors(nums)
		fmt.Println(nums)
	}

}
