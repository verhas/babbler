// Command id-tool is a small CLI wrapper around the identifier package.
package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/verhas/babbler/identifier"
)

func main() {
	if len(os.Args) != 3 || os.Args[1] != "encode" {
		fmt.Fprintln(os.Stderr, "usage: id-tool encode <num>")
		os.Exit(1)
	}

	num, err := strconv.Atoi(os.Args[2])
	if err != nil {
		fmt.Fprintf(os.Stderr, "invalid number: %v\n", err)
		os.Exit(1)
	}

	id, err := identifier.NumberToID(num)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}

	fmt.Println(id)
}
