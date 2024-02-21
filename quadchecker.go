package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func runCommand(fileName string, args ...string) string {
	cmd := exec.Command(fileName, args...)
	output, _ := cmd.CombinedOutput()
	return string(output)
}

func isQuadFunction(fileName, x, y, str string) bool {
	return str == runCommand(fileName, x, y)
}

func findMatchingFiles(x, y int, str string) []string {
	files := []string{"./quadA", "./quadB", "./quadC", "./quadD", "./quadE"}
	var results []string

	for _, file := range files {
		if isQuadFunction(file, strconv.Itoa(x), strconv.Itoa(y), str) {
			results = append(results, file)
		}
	}

	return results
}

func QuadChecker(str string) {
	var resultStrings []string

	y, x := 0, 0
	endX := false

	for _, char := range str {
		if char == '\n' {
			endX = true
			y++
		}

		if !endX {
			x++
		}
	}

	matchingFiles := findMatchingFiles(x, y, str)

	if len(matchingFiles) < 1 || str == "" {
		fmt.Println("Not a quad function")
		return
	}

	for i, file := range matchingFiles {
		resultStrings = append(resultStrings,
			fmt.Sprintf("[%v] [%v] [%v]", file[2:], x, y))

		if i < len(matchingFiles)-1 {
			resultStrings = append(resultStrings, " || ")
		}
	}

	fmt.Println(strings.Join(resultStrings, ""))
}

func main() {
	args := os.Args[1:]
	QuadChecker(args[0])
}
