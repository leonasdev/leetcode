package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	number := os.Args[1]
	for _, ch := range number {
		if !unicode.IsDigit(ch) {
			fmt.Fprintln(os.Stdout, "First arg must be number")
			os.Exit(1)
		}
	}

	_, err := strconv.Atoi(number)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	fixNumber := ""
	for i := 0; i < 4-len(number); i++ {
		fixNumber += "0"
	}
	fixNumber += number

	name := os.Args[2]
	name = strings.ReplaceAll(name, "-", "_")

	directory := fixNumber + "_" + name
	if err := os.Mkdir(directory, 0755); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	file, err := os.Create(directory + "/" + directory + ".go")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer file.Close()

	fmt.Fprintf(file, "package leetcode")

	testFile, err := os.Create(directory + "/" + directory + "_test.go")
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
	defer testFile.Close()

	fmt.Fprintf(testFile, "package leetcode")

	fmt.Println("Successfully create " + directory)
}
