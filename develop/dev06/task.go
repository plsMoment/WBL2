package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"slices"
	"strconv"
	"strings"
)

type fieldsFlag []int

func (f *fieldsFlag) String() string {
	return fmt.Sprint(*f)
}

func (f *fieldsFlag) Set(s string) error {
	numStr := strings.Split(s, ",")
	for _, str := range numStr {
		num, err := strconv.Atoi(str)
		if err != nil {
			return err
		}
		*f = append(*f, num)
	}
	slices.Sort(*f)
	*f = slices.Compact(*f)
	return nil
}

type parsedFlags struct {
	delimitedOnly bool
	delimiter     string
	fields        fieldsFlag
}

var flags = parsedFlags{}

// readStdin function reads text from stdin and returns slice of strings
func readStdin() []string {
	var data []string

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		data = append(data, scanner.Text())
	}

	if scanner.Err() != nil {
		log.Fatal("reading from stdin failed")
	}
	return data
}

// cut function cut strings, result data depends on flags from struct parsedFlags
func cut() []string {
	var res []string

	data := readStdin()
	for _, str := range data {
		delimitedStr := strings.Split(str, flags.delimiter)

		if len(delimitedStr) < 2 {
			if flags.delimitedOnly {
				continue
			} else {
				res = append(res, str)
				continue
			}
		}

		var sb strings.Builder
		for _, num := range flags.fields {
			if num <= len(delimitedStr) && num > 0 {
				sb.WriteString(delimitedStr[num-1])
				sb.WriteString(flags.delimiter)
			}
		}
		res = append(res, strings.TrimSuffix(sb.String(), flags.delimiter))
	}
	return res
}

func main() {
	flag.StringVar(&flags.delimiter, "d", "\t", "delimiter instead of TAB")
	flag.BoolVar(
		&flags.delimitedOnly, "s", false,
		"if this option encountered, cut does not output lines where there is no separator",
	)
	flag.Var(&flags.fields, "f", "number of fields to cut")
	flag.Parse()

	if len(flags.fields) == 0 {
		log.Fatal("you must use -f parameter")
	}

	delimitedText := cut()
	for _, str := range delimitedText {
		fmt.Println(str)
	}
}
