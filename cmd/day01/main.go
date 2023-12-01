package main

import (
	"flag"
	"fmt"
	"os"
)

func main() {
	fn := parseFlags()
	fmt.Println(fn)
}

func parseFlags() string {
	fn := flag.String("f", "", "Path to your puzzle input file")
	flag.Parse()

	if *fn == "" {
		fmt.Println("Usage: program -f <filename>")
		os.Exit(1)
	}

	return *fn
}
