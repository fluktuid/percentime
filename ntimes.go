package main

import (
	"bytes"
	"io"
	"os/exec"
	"strconv"
	"strings"
	"sync"
	"time"
)

func Ntimes(cnt int, cmdName string, cmdArgs []string, stdin io.Reader, stderr io.Writer, stdoutCh chan float64, parallels int) {
	var wg sync.WaitGroup

	sema := make(chan bool, parallels)

	for i := 0; i < cnt; i++ {
		wg.Add(1)
		go func() {
			sema <- true

			defer func() {
				wg.Done()
				<-sema
			}()

			stdoutBuffer := new(bytes.Buffer)

			cmd := exec.Command(cmdName, cmdArgs...)
			if opts.UseCmdResp {
				cmd.Stdin = stdin
				cmd.Stdout = stdoutBuffer
				cmd.Stderr = nilWriter{}
				err := cmd.Run()

				if err != nil {
					panic(err)
				}

				f, err := read(stdoutBuffer)

				if err != nil {
					panic(err)
				}

				stdoutCh <- f
			} else {
				cmd.Stdin = nilReader{}
				cmd.Stdout = nilWriter{}
				cmd.Stderr = nilWriter{}
				start := time.Now()
				err := cmd.Run()
				elapsed := timeTrack(start)

				if err != nil {
					panic(err)
				}
				stdoutCh <- elapsed.Seconds()
			}
		}()
	}

	wg.Wait()
}

func timeTrack(start time.Time) time.Duration {
	elapsed := time.Since(start)
	return elapsed
}

func read(buff *bytes.Buffer) (float64, error) {
	str := buff.String()

	lineWithDot := strings.Replace(str, ",", ".", -1)

	f, convErr := strconv.ParseFloat(lineWithDot, 64)
	if convErr != nil {
		return 0, convErr
	}
	return f, nil
}
