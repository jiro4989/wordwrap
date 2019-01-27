package wordwrap

import (
	"strings"
)

func WordWrap(text string, n int) []string {
	var ret []string
	for _, t := range strings.Split(text, "\n") {
		var wordLen int
		var current []string
		var arr []string
		for _, c := range t {
			if IsHalfChar(c) {
				wordLen++
			} else {
				wordLen += 2
			}
			if c == ' ' || c == '　' || c == '\t' {
				current = append(current, strings.Join(arr, ""))
				arr = []string{}
				continue
			}
			if n < wordLen {
				ret = append(ret, strings.Join(current, ""))
				current = []string{}
				wordLen = 0
			}
			arr = append(arr, string(c))
		}
		if 0 < len(arr) {
			ret = append(ret, strings.Join(arr, ""))
		}
	}
	return ret
}

// WordWrapFixedWidth は固定長で文字列を折り返す。
func WordWrapFixedWidth(text string, n int) []string {
	ret := []string{}
	for _, line := range strings.Split(text, "\n") {
		// すでに文字列の長さがn未満ならその行の計算はスキップ
		if len(line) < n {
			continue
		}

		var cnt int
		var buf []rune
		for _, r := range line {
			if IsHalfChar(r) {
				cnt++
			} else {
				cnt += 2
			}
			buf = append(buf, r)
			if n <= cnt {
				ret = append(ret, string(buf))
				buf = []rune{}
				cnt = 0
			}
		}
		if 0 < len(buf) {
			ret = append(ret, string(buf))
		}
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
