package stringutil

func ReverseRange(s string) string {
	var outLen int = len(s)
	var out []rune = make([]rune, outLen)

	for index, runeValue := range s {
		out[outLen-index-1] = runeValue
	}

	return string(out)
}

func ReverseConvert(s string) string {
	var out []rune = []rune(s)
	var outLen int = len(out)

	// i: 17 / 2 = 8 (1)
	// j: 17-8-1 = 8

	// i: 18 / 2 = 9 (0)
	// j: 18-9-1 = 8
	for i := outLen / 2; i < outLen; i++ {
		j := outLen - i - 1
		out[i], out[j] = out[j], out[i]
	}

	return string(out)
}
