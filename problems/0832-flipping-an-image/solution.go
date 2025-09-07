package flippinganimage

func flipAndInvertImage(image [][]int) [][]int {
	for _, row := range image {
		reverse(row)
		invert(row)
	}
	return image
}

func reverse(row []int) {
	l, r := 0, len(row)-1
	for l < r {
		row[l], row[r] = row[r], row[l]
		l++
		r--
	}
}

func invert(row []int) {
	for i := 0; i < len(row); i++ {
		if row[i] == 1 {
			row[i] = 0
		} else {
			row[i] = 1
		}
	}
}
