package main

import (
	"os"
	"regexp"
	"strings"
)

func main() {
	dirs, err := os.ReadDir(".")
	if err != nil {
		panic(err)
	}
	for _, dir := range dirs {
		re, err := regexp.Compile(`^[0-9]*_`)
		if err != nil {
			panic(err)
		}
		numLen := len(re.FindString(dir.Name())) - 1
		if numLen > 0 && numLen < 4 {
			err := os.Rename(dir.Name(), strings.Repeat("0", 4-numLen)+dir.Name())
			if err != nil {
				panic(err)
			}
		}
	}
}
