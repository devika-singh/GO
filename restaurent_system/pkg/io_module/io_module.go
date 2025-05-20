package io_module

import (
	"bufio"
	"fmt"
	"strings"
)

func GetInput(prompt string, r *bufio.Reader) (string, error) {
	fmt.Print(prompt)
	input, err := r.ReadString('\n')

	return strings.TrimSpace(input), err
}

func ClearScreen() { // just print two ANSI escapes
	fmt.Print("\033[2J\033[H") // 2J = erase; H = cursor-home
}
