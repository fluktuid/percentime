package main

import (
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"

	"github.com/jessevdk/go-flags"
)

const AppName = "percentime"

type Options struct {
	Parallels   int  `short:"p" long:"parallels" description:"Parallel degree of execution" default:"1"`
	ShowVersion bool `short:"v" long:"version" description:"Show version"`
	UseCmdResp  bool `short:"c" long:"commandResponse" description:"Use command Response instead of measured time"`
}

var opts Options

var numbers sort.Float64Slice

var values = [9]int{50, 66, 75, 80, 90, 95, 98, 99, 100}

func main() {
	parser := flags.NewParser(&opts, flags.Default^flags.PrintErrors)
	parser.Name = AppName
	parser.Usage = "N [OPTIONS] -- COMMAND"

	args, err := parser.Parse()

	if err != nil {
		_, _ = fmt.Fprint(os.Stderr, err)
		return
	}

	if len(args) == 0 {
		parser.WriteHelp(os.Stdout)
		return
	}

	if opts.ShowVersion {
		_, _ = io.WriteString(os.Stdout, fmt.Sprintf("%s v%s, build %s\n", AppName, Version, GitCommit))
		return
	}

	cnt, err := strconv.Atoi(args[0])
	if err != nil {
		panic(err)
	}

	cmdName := args[1]
	cmdArgs := args[2:]

	stdoutCh := make(chan float64)
	exitCh := make(chan bool)

	go wrapper(os.Stdout, stdoutCh, exitCh)

	Ntimes(cnt, cmdName, cmdArgs, os.Stdin, os.Stderr, stdoutCh, opts.Parallels)
	exitCh <- true
}

func wrapper(stdout io.Writer, stdoutCh chan float64, exitCh chan bool) {
	for {
		select {
		case r := <-stdoutCh:
			err := percentime(r, stdout, opts)
			if err != nil {
				panic(err)
			}
		case <-exitCh:
			return
		}
	}
}
