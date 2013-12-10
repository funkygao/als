package als

import (
	"fmt"
)

func GroupedLevel(level int) string {
	return groupInts(level, GROUP_LEVELS)
}

func GroupedXP(xp int) string {
	return groupInts(xp, GROUP_XP)
}

func groupInts(i int, groups []int) string {
	prev, current := 0, 0
	for _, v := range groups {
		if v >= i {
			current = v
			break
		}

		prev = v
	}

	return fmt.Sprintf("%d-%d", prev, current)
}
