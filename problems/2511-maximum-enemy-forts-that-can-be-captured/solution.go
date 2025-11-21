package maximumenemyfortsthatcanbecaptured

func captureForts(forts []int) int {
	var lastidx, lastfort, maxdist int

	for i, f := range forts {
		if f != 0 {
			if f == -lastfort {
				maxdist = max(maxdist, i-lastidx-1)
			}
			lastfort = f
			lastidx = i
		}
	}

	return maxdist
}
