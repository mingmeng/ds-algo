package string

import "strings"

// replaceWords 暴力解法 hash +
func replaceWords(dictionary []string, sentence string) string {
	words := strings.Split(sentence, " ")
	if len(words) == 0 {
		return sentence
	}
	dicMap := make(map[string]bool, len(dictionary))
	for _, w := range dictionary {
		dicMap[w] = true
	}

	for index, s := range words {
		for i := 0; i < len(s); i++ {
			pre := s[0:i]
			if _, ok := dicMap[pre]; ok {
				words[index] = pre
				break
			}
		}
	}

	return strings.Join(words, " ")
}
