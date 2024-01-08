package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

type parsedFlags struct {
	c, i, v, f, n bool
	a, b, ctx     int
	pattern, file string
}

var flags = parsedFlags{}

// grepInt needs to optimize work when number of matches required
func grepInt() int {
	in, err := os.Open(flags.file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := in.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	elemsNum := 0
	scanner := bufio.NewScanner(in)
	if flags.f {
		if flags.i {
			flags.pattern = strings.ToLower(flags.pattern)
			for scanner.Scan() {
				elemsNum += strings.Count(strings.ToLower(scanner.Text()), flags.pattern)
			}
		} else {
			for scanner.Scan() {
				elemsNum += strings.Count(scanner.Text(), flags.pattern)
			}
		}
	} else {
		var r *regexp.Regexp
		if flags.i {
			r = regexp.MustCompile("(?i)" + flags.pattern)
		} else {
			r = regexp.MustCompile(flags.pattern)
		}
		for scanner.Scan() {
			elemsNum += len(r.FindAllIndex(scanner.Bytes(), -1))
		}
	}
	return elemsNum
}

// grepStr returns array of strings where pattern was found with parameter in flags
func grepStr() []string {
	in, err := os.Open(flags.file)
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := in.Close(); err != nil {
			log.Fatal(err)
		}
	}()

	var rowsIndexes []int // contains number of rows where string matches with pattern
	scanner := bufio.NewScanner(in)
	row := 1
	if flags.f {
		if flags.i {
			flags.pattern = strings.ToLower(flags.pattern)
			for scanner.Scan() {
				if strings.Contains(strings.ToLower(scanner.Text()), flags.pattern) {
					rowsIndexes = append(rowsIndexes, row)
				}
				row++
			}
		} else {
			for scanner.Scan() {
				if strings.Contains(scanner.Text(), flags.pattern) {
					rowsIndexes = append(rowsIndexes, row)
				}
				row++
			}
		}
	} else {
		var r *regexp.Regexp
		if flags.i {
			r = regexp.MustCompile("(?i)" + flags.pattern)
		} else {
			r = regexp.MustCompile(flags.pattern)
		}
		for scanner.Scan() {
			if r.Match(scanner.Bytes()) {
				rowsIndexes = append(rowsIndexes, row)
			}
			row++
		}
	}

	if len(rowsIndexes) < 1 {
		return nil
	}

	if _, err = in.Seek(0, 0); err != nil {
		log.Fatal(err)
	}

	var res []string
	scanner = bufio.NewScanner(in)
	rowArrIdx := 0

	if flags.v {
		for i := 1; scanner.Scan(); i++ {
			if rowArrIdx < len(rowsIndexes) && i == rowsIndexes[rowArrIdx] {
				rowArrIdx++
				continue
			}
			res = append(res, fmt.Sprintf("%d:%s", i, scanner.Text()))
		}
	} else {
		lastRowIndex := rowsIndexes[len(rowsIndexes)-1] + flags.a
		for i := 1; i <= lastRowIndex; i++ {
			if !scanner.Scan() {
				break
			}

			if i >= rowsIndexes[rowArrIdx]-flags.b && i <= rowsIndexes[rowArrIdx]+flags.a {
				if i == rowsIndexes[rowArrIdx] {
					res = append(res, fmt.Sprintf("%d:%s", i, scanner.Text()))
				} else {
					res = append(res, fmt.Sprintf("%d-%s", i, scanner.Text()))
				}

				if i == rowsIndexes[rowArrIdx]+flags.a {
					rowArrIdx++
				}
			}
		}
	}

	if !flags.n {
		for i, str := range res {
			res[i] = str[strings.IndexAny(str, ":-")+1:]
		}
	}

	return res
}

func main() {
	parseFlags()

	if flags.c {
		fmt.Println(grepInt())
	} else {
		foundText := grepStr()
		for _, str := range foundText {
			fmt.Println(str)
		}
	}
}

func parseFlags() {
	flag.BoolVar(&flags.c, "c", false, "print the number of rows found")
	flag.BoolVar(&flags.i, "i", false, "ignore the case of characters")
	flag.BoolVar(&flags.v, "v", false, "print only those lines which was not found")
	flag.BoolVar(&flags.f, "F", false, "search pattern as a regular string, not a regular expression")
	flag.BoolVar(&flags.n, "n", false, "print numbers of strings where match found")
	flag.IntVar(&flags.a, "A", 0, "show the match and n lines after it")
	flag.IntVar(&flags.b, "B", 0, "show the match and n lines before it")
	flag.IntVar(&flags.ctx, "C", 0, "show the match and n lines before and after it")
	flag.Parse()

	if flags.a == 0 {
		flags.a = flags.ctx
	}
	if flags.b == 0 {
		flags.b = flags.ctx
	}

	if len(flag.Args()) > 1 {
		flags.pattern = flag.Arg(0)
		flags.file = flag.Arg(1)
	} else {
		log.Fatal("not enough arguments, enter pattern then filepath")
	}
}
