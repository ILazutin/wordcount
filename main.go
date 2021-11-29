package main

import (
	"errors"
	"flag"
	"fmt"
	"strings"
)

func main() {
	src, err := readInput()
	if err != nil {
		fail(err)
	}
	fmt.Println(wordcount(src))
}

func wordcount(src string) int {
	return len(strings.Fields(src))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string, err error) {
	//flag.StringVar(&pattern, "p", "", "pattern to match against")
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	if src == "" {
		return src, errors.New("missing string to match")
	}
	return src, nil
}

// fail prints the error and exits.
func fail(err error) {
	fmt.Println("match:", err)
	os.Exit(1)
}
