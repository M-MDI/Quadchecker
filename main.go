package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	var stdInInput [5000]byte
	n, _ := os.Stdin.Read(stdInInput[:])
	str := string(stdInInput[:n])

	cmd := exec.Command("./quadchecker", str)

	output, _ := cmd.CombinedOutput()

	fmt.Print(string(output))
}
