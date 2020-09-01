package main

import (
	"fmt"
	"os"
	"path"
	"strings"

	"github.com/snarechops/assembler/parse"
)

func check(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	// Process cli args
	args := processArgs(os.Args)

	// Load file
	filename := path.Base(args.Input)
	basename := strings.Replace(filename, path.Ext(args.Input), "", -1)
	outputname := basename + ".bin"

	file, err := os.Open(args.Input)
	check(err)
	defer file.Close()

	// Parse file
	data, err := parse.File(file)
	check(err)

	// Output binary
	output, err := os.Create(outputname)
	check(err)
	defer output.Close()

	size, err := output.Write(data)
	check(err)

	fmt.Printf("Wrote %d bytes to %s", size, outputname)
}

type options struct {
	Input string
}

func processArgs(args []string) *options {

	return &options{
		Input: path.Clean(args[1]),
	}
}
