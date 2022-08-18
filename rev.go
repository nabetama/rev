package main

import (
	"flag"
	"fmt"
	"io"
)

const (
	ExitCodeOK             = 0
	ExitCodeParseFlagError = 1
)

type CLI struct {
	outStream, errStream io.Writer
}

func rev(s string) (ret string) {
	for _, v := range s {
		ret = string(v) + ret
	}
	return ret
}

func (c *CLI) Run(args []string) int {
	flags := flag.NewFlagSet("rev", flag.ContinueOnError)
	flags.SetOutput(c.errStream)

	if err := flags.Parse(args[1:]); err != nil {
		return ExitCodeParseFlagError
	}

	fmt.Fprintf(c.outStream, rev(args[1])+"\n")
	return ExitCodeOK
}
