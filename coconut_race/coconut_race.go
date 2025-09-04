package coconut_race

import "slices"

func canReach(leftPos, rightPos, leftSpeed, rightSpeed, target int) bool {
	if leftPos == rightPos {
		return true
	}
	if leftSpeed <= rightSpeed {
		return false
	}
	approachSpeed := leftSpeed - rightSpeed
	approachDist := rightPos - leftPos
	rightDist := target - rightPos
	return approachDist*rightSpeed <= approachSpeed*rightDist
}

type Coconut struct {
	pos   int
	speed int
}

func countCoconutBouquets(target int, position []int, speed []int) int {
	if len(position) == 0 {
		return 0
	}
	coconuts := make([]Coconut, 0, len(position))
	for i, pos := range position {
		coconuts = append(coconuts, Coconut{pos: pos, speed: speed[i]})
	}
	slices.SortFunc(coconuts, func(a, b Coconut) int { return a.pos - b.pos })

	lastInd := len(coconuts) - 1
	bouquetPos := coconuts[lastInd].pos
	bouquetSpeed := coconuts[lastInd].speed
	res := 0
	for i := len(coconuts) - 2; i >= 0; i-- {
		curPos := coconuts[i].pos
		curSpeed := coconuts[i].speed
		if canReach(curPos, bouquetPos, curSpeed, bouquetSpeed, target) {
			continue
		}
		res++
		bouquetPos = curPos
		bouquetSpeed = curSpeed
	}
	return res + 1
}
