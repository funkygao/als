package als

import (
	"fmt"
)

func GroupedLevel(level int) string {
	prev, current := 1, 1
	for _, lv := range GROUP_LEVELS {
		if lv >= level {
			current = lv
			break
		}

		prev = lv
	}

	return fmt.Sprintf("%d-%d", prev, current)
}
