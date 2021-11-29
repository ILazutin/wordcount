package main

import (
	"flag"
	"fmt"
	"strings"
)

func main() {
	src := readInput()
	fmt.Println(wordcount(src))
}

func wordcount(src string) int {
	return len(strings.Fields(src))
}

// readInput reads pattern and source string
// from command line arguments and returns them.
func readInput() (src string) {
	//flag.StringVar(&pattern, "p", "", "pattern to match against")
	flag.Parse()
	src = strings.Join(flag.Args(), "")
	return src
}
