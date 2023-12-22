package leetcode

import (
	"strconv"
	"strings"
)

type Codec struct {
	b strings.Builder
}

func (codec *Codec) Encode(strs []string) string {
	for _, s := range strs {
		codec.b.WriteString(strconv.Itoa(len(s)))
		codec.b.WriteRune('#')
		codec.b.WriteString(s)
	}

	return codec.b.String()
}

func (codec *Codec) Decode(str string) []string {
	res := []string{}
	i := 0
	for i < len(str) {
		j := i
		for str[j] != '#' {
			j++
		}
		num, _ := strconv.Atoi(str[i:j])
		res = append(res, str[j+1:j+1+num])
		i = j + 1 + num
	}

	return res
}
