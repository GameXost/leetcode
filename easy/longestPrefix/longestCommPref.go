package prefix

func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	minLen := len(strs[0])

	for i := 1; i < len(strs); i++ {
		curWord := strs[i-1]
		nextWord := strs[i]
		if len(nextWord) < minLen {
			minLen = len(nextWord)
		}
		j := 0
		for j < minLen && curWord[j] == nextWord[j] {
			j++
		}
		minLen = j
		if minLen == 0 {
			break
		}

	}
	return strs[0][:minLen]
}
