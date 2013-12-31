package als

func Colorize(colors []string, str string) string {
	r := ""
	for _, color := range colors {
		r += color_table[color]
	}

	r += str + color_table["Reset"]
	return r
}
