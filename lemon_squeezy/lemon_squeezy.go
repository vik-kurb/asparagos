package lemon_squeezy

import (
	"cmp"
	"slices"
)

type Lemon struct {
	iq  int
	age int
}

func findBestTeamIQ(iqs []int, ages []int) int {
	lemons := make([]Lemon, len(iqs))
	for i, iq := range iqs {
		lemons[i] = Lemon{iq: iq, age: ages[i]}
	}
	slices.SortFunc(lemons, func(a, b Lemon) int {
		if a.age == b.age {
			return cmp.Compare(a.iq, b.iq)
		}
		return cmp.Compare(a.age, b.age)
	})

	teamIQs := make([]int, len(lemons))
	maxTeamIQ := 0
	for i, lemon := range lemons {
		teamIQs[i] = lemon.iq
		for prevInd := 0; prevInd < i; prevInd++ {
			if lemons[prevInd].iq > lemon.iq {
				continue
			}
			newIQ := teamIQs[prevInd] + lemon.iq
			if newIQ > teamIQs[i] {
				teamIQs[i] = newIQ
			}
		}
		if teamIQs[i] > maxTeamIQ {
			maxTeamIQ = teamIQs[i]
		}
	}
	return maxTeamIQ
}
