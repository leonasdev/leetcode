package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
)

// Rewrite all package from "main" to "leetcode"
func main() {
	cw, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dirs, err := os.ReadDir(cw)
	if err != nil {
		panic(err)
	}

	for _, dir := range dirs {
		if !dir.IsDir() || dir.Name() == ".git" {
			continue
		}
		subDirs, err := os.ReadDir(dir.Name())
		if err != nil {
			panic(err)
		}
		for _, subDir := range subDirs {
			if subDir.IsDir() {
				continue
			}
			reg := regexp.MustCompile(`^` + dir.Name())
			if reg.MatchString(subDir.Name()) {
				f, err := os.ReadFile(filepath.Join(dir.Name(), subDir.Name()))
				if err != nil {
					panic(err)
				}
				newString := strings.ReplaceAll(string(f), "package main", "package leetcode")
				fmt.Println(string(f))
				fmt.Println(newString)
			}
		}
	}
}
