// ex 1.4
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	names := make(map[string][]string)

	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, names)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, names)
			f.Close()
		}
	}

	for line, n := range counts {
		if n > 1 {
			fmt.Printf("\"%s\" occurs %d times from %v file(s)).\n", line, n, names[line])
		}
	}
}

func in(needle string, strings []string) bool {
	for _, s := range strings {
		if needle == s {
			return true
		}
	}
	return false
}

func countLines(f *os.File, counts map[string]int, names map[string][]string) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		counts[line]++
		if !in(f.Name(), names[line]) {
			names[line] = append(names[line], f.Name())
		}
	}
}
