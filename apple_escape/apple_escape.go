package apple_escape

import (
	"cmp"
	"slices"
)

type Trip struct {
	from int
	to   int
	num  int
}

type Event struct {
	location int
	num      int
	isFrom   bool
}

func canSaveAllApples(trips []Trip, capacity int) bool {
	events := make([]Event, 0, 2*len(trips))
	for _, trip := range trips {
		events = append(events, Event{trip.from, trip.num, true}, Event{trip.to, trip.num, false})
	}
	slices.SortFunc(events, func(a, b Event) int {
		if a.location != b.location {
			return cmp.Compare(a.location, b.location)
		}
		if a.isFrom == b.isFrom {
			return 0
		}
		if !a.isFrom {
			return -1
		}
		return 1
	})
	num := 0
	for _, event := range events {
		if !event.isFrom {
			num -= event.num
			continue
		}
		if capacity-num < event.num {
			return false
		}
		num += event.num
	}
	return true
}
