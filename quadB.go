package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func QuadB(x, y int) {
	if x > 0 && y > 0 {
		for m := 1; m <= y; m++ {
			for b := 1; b <= x; b++ {
				if m == 1 && b == 1 {
					z01.PrintRune('/')
				} else if m == 1 && b > 1 && b < x {
					z01.PrintRune('*')
				} else if m == 1 && b == x {
					z01.PrintRune('\\')
				} else if b == 1 && m > 1 && m < y {
					z01.PrintRune('*')
				} else if m > 1 && m < y && b > 1 && b < x {
					z01.PrintRune(' ')
				} else if m > 1 && m < y && b == x {
					z01.PrintRune('*')
				} else if m == y && b == 1 {
					z01.PrintRune('\\')
				} else if m == y && b > 1 && b < x {
					z01.PrintRune('*')
				} else {
					z01.PrintRune('/')
				}
			}
			z01.PrintRune('\n')
		}
	}
}

func main() {
	args := os.Args[1:]
	x, _ := strconv.Atoi(args[0])
	y, _ := strconv.Atoi(args[1])

	QuadB(x, y)
}
