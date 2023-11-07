package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// Rename all main.go to the format: 0123_the_leetcode_question.go
func main() {
	cw, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	dirs, err := os.ReadDir(cw)
	if err != nil {
		panic(err)
	}

	reg := regexp.MustCompile(`^main`)
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
			if reg.MatchString(subDir.Name()) {
				oldDirName := dir.Name() + "/" + subDir.Name()
				newDirName := strings.ReplaceAll(oldDirName, "main", dir.Name())
				fmt.Println(oldDirName + "->" + newDirName)
				if err := os.Rename(oldDirName, newDirName); err != nil {
					panic(err)
				}
			}
		}
	}
}
