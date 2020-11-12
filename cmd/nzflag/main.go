package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/zakuro9715/nzflag"
)

func main() {
	args := nzflag.New().NormalizeArgsToStrings()
	fmt.Printf("IN : %v\n", strings.Join(os.Args[1:], " "))
	fmt.Printf("OUT: %v\n", strings.Join(args, " "))
}
