package main

import (
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
				targetFile := filepath.Join(dir.Name(), subDir.Name())
				f, err := os.ReadFile(targetFile)
				if err != nil {
					panic(err)
				}
				newContent := []byte(strings.ReplaceAll(string(f), "package main", "package leetcode"))
				os.WriteFile(targetFile, newContent, 0644)
			}
		}
	}
}
