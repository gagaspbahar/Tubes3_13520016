package stringmatching


func CalculatePrefixSuffix(substr string) []int {
	table := make([]int, len(substr))
	j := 0
	i := 1
	for i < len(substr) {
		if substr[i] == substr[j] {
			table[i] = j + 1
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = table[j]
		}
	}
	return table
}

func KMP(str string, substr string) int {
	table := CalculatePrefixSuffix(substr)
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
	i, j := 0, 0
	for i < len_str {
		if str[i] == substr[j] {
			i++
			j++
		} else if j == 0 {
			i++
		} else {
			j = table[j-1]
		}
		if len_sub == j {
			return 1
		}
	}
	return 0
}