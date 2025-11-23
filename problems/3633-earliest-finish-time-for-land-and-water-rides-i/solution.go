package earliestfinishtimeforlandandwaterridesi

func earliestFinishTime(landStartTime []int, landDuration []int, waterStartTime []int, waterDuration []int) int {
	n, m := len(landStartTime), len(waterStartTime)
	land, water := 1<<31, 1<<31
	land_water, water_land := 1<<31, 1<<31

	for i := 0; i < n; i++ {
		land = min(land, landStartTime[i]+landDuration[i])
	}

	for i := 0; i < m; i++ {
		water = min(water, waterStartTime[i]+waterDuration[i])
		land_water = min(land_water, max(waterStartTime[i], land)+waterDuration[i])
	}

	for i := 0; i < n; i++ {
		water_land = min(water_land, max(landStartTime[i], water)+landDuration[i])
	}
	return min(water_land, land_water)
}
