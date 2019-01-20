package wordwrap

import "strings"

func WordWrap(text string, n int) []string {
	max := n + 10
	var wordLen int
	var arrIdx int
	var ret []string
	arr := make([]string, max)
	for _, c := range text {
		if IsHalfChar(c) {
			wordLen++
		} else {
			wordLen += 2
		}

		if n < wordLen {
			ret = append(ret, strings.TrimSpace(strings.Join(arr, "")))
			arr = make([]string, max)
			arrIdx = 0
			wordLen = 0
		}
		arrIdx++
		arr[arrIdx] = string(c)
	}
	if 0 < wordLen {
		ret = append(ret, strings.TrimSpace(strings.Join(arr, "")))
	}
	return ret
}

func IsHalfChar(r rune) bool {
	var i uint
	for i = 0x00; i <= 0x7f; i++ {
		if r == rune(i) {
			return true
		}
	}
	return false
}
