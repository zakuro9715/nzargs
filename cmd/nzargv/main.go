package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zakuro9715/nzargv"
)

func main() {
	args, err := nzargv.New().NormalizeArgsToStrings()
	if err != nil {
		fmt.Fprintln(os.Stdout, err)
		os.Exit(1)
	}
	fmt.Printf("IN : %v\n", strings.Join(os.Args[1:], " "))
	fmt.Printf("OUT: %v\n", strings.Join(args, " "))

}
