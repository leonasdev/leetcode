package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
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
				source, err := os.ReadFile(targetFile)
				if err != nil {
					panic(err)
				}
				if err := os.WriteFile(targetFile[:len(targetFile)-3]+"_test.go", source, 0644); err != nil {
					panic(err)
				}
				fmt.Printf("Add file: %s\n", targetFile)
			}
		}
	}
}
