package strawberry_fever

func canEatAllStrawberries(storeStock []int, wayCost []int) int {
	minEnergyReserve := 0
	minEnergyReserveInd := 0
	energyReserve := 0
	for ind := 0; ind < len(storeStock); ind++ {
		if energyReserve < minEnergyReserve {
			minEnergyReserve = energyReserve
			minEnergyReserveInd = ind
		}
		energyReserve += storeStock[ind] - wayCost[ind]
	}
	if energyReserve < 0 {
		return -1
	}
	return minEnergyReserveInd
}
