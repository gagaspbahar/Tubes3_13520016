package boyermoore

func Boyermoore(str string, substr string) int {
	table := CalculateTable(substr)
	len_sub := len(substr)
	len_str := len(str)
	switch {
	case len_sub > len_str:
		return 0
	case len_sub == len_str:
		if str == substr {
			return 1
		}
		return 0
	}
	i := 0
	for i - 1 < len_str - len_sub {
		j := len_sub - 1
		for (j >= 0 && str[i+j] == substr[j]) {
			j--
		}
		if j < 0 {
			return 1
		}
		slide := j - table[str[i+j]]
		if slide < 1 {
			slide = 1
		}
		i += slide
	}
	return 0
}

func CalculateTable(substr string) [256]int {
	var table [256]int
	for i := 0; i < 256; i++ {
		table[i] = -1
	}
	for i := 0; i < len(substr); i++ {
		table[substr[i]] = i
	}
	return table
}