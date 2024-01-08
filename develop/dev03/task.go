package main

import (
	"bufio"
	"cmp"
	"flag"
	"fmt"
	"log"
	"math/big"
	"os"
	"slices"
	"strings"
)

type parsedFlags struct {
	n, r, u  bool
	k        int
	filepath string
}

var flags = parsedFlags{}

// readFile split text in words
func readFile() [][]string {
	in, err := os.Open(flags.filepath)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := in.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	scanner := bufio.NewScanner(in)
	var text [][]string

	if flags.u {
		set := make(map[string]int)
		var lines []string

		for scanner.Scan() {
			line := scanner.Text()
			set[line] += 1
			lines = append(lines, line)
		}
		lines = slices.DeleteFunc(lines, func(s string) bool {
			return set[s] > 1
		})

		for _, line := range lines {
			text = append(text, strings.Split(line, " "))
		}
	} else {
		for scanner.Scan() {
			text = append(text, strings.Split(scanner.Text(), " "))
		}
	}

	return text
}

// sort text
func sort(text [][]string) {
	if flags.n {
		slices.SortFunc(text, func(a, b []string) int {
			if flags.k >= len(a) {
				return -1
			}
			if flags.k >= len(b) {
				return 1
			}
			n1 := new(big.Int)
			n2 := new(big.Int)
			n1.SetString(a[flags.k], 10)
			n2.SetString(b[flags.k], 10)
			return n1.Cmp(n2)
		})
	} else {
		slices.SortFunc(text, func(a, b []string) int {
			if flags.k >= len(a) {
				return -1
			}
			if flags.k >= len(b) {
				return 1
			}
			return cmp.Compare(a[flags.k], b[flags.k])
		})
	}
}

func main() {
	parseFlags()
	text := readFile()
	sort(text)
	if flags.r {
		slices.Reverse(text)
	}

	for _, line := range text {
		fmt.Println(strings.Join(line, " "))
	}
}

func parseFlags() {
	flag.BoolVar(&flags.n, "n", false, "sort by numeric value")
	flag.BoolVar(&flags.r, "r", false, "sort in reverse order")
	flag.BoolVar(&flags.u, "u", false, "ignore duplicate lines")
	flag.IntVar(&flags.k, "k", 1, "specifying the column to sort")
	flag.Parse()

	if flags.k > 0 {
		flags.k -= 1
	} else {
		log.Fatal("-k value must be positive")
	}

	if len(flag.Args()) > 0 {
		flags.filepath = flag.Arg(0)
	} else {
		log.Fatal("not enough arguments, enter filepath")
	}
}
