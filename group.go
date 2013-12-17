package als

import (
	"fmt"
)

func GroupedLevel(level int) string {
	return GroupInt(level, GROUP_LEVELS)
}

func GroupedXP(xp int) string {
	return GroupInt(xp, GROUP_XP)
}

func GroupedSessionLen(sessionLen int) string {
	return GroupInt(sessionLen, GROUP_SESSIONLEN)
}

func LevelLabels() []string {
	return intsGroupLabels(GROUP_LEVELS)
}

func SessionLenLabels() []string {
	return intsGroupLabels(GROUP_SESSIONLEN)
}

func intsGroupLabels(group []int) []string {
	ret := make([]string, 0)
	n := len(group)
	for idx, val := range group {
		if idx == n-1 {
			break
		}
		ret = append(ret, fmt.Sprintf("%d-%d", val, group[idx+1]))
	}
	return ret

}

func GroupInt(i int, groups []int) string {
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

