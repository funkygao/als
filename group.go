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

// [x, y)
func GroupInt(val int, groups []int) string {
	const (
		INVALID_LEFT    = -109
		INVALID_CURRENT = -107
	)
	if len(groups) <= 2 {
		return ""
	}

	left, current := INVALID_LEFT, INVALID_CURRENT
	for _, needle := range groups {
		if val < needle {
			current = needle
			break
		}

		left = needle
	}

	if left == INVALID_LEFT || current == INVALID_CURRENT {
		return ""
	}

	return fmt.Sprintf("%d-%d", left, current)
}
