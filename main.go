package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"os"
)

var (
	fileName        string
	reader          io.Reader
	readingFromFile bool
	isByteMode      bool
	isLineMode      bool
	isWordMode      bool
)

func init() {
	flag.Usage = func() {
		fmt.Fprintf(os.Stderr, "Usage: %s [OPTIONS] [FILE]\n\n", os.Args[0])
		fmt.Fprintf(os.Stderr, "Options:\n")
		fmt.Fprintf(os.Stderr, "  -c\tto count bytes\n")
		fmt.Fprintf(os.Stderr, "  -l\tto count lines\n")
		fmt.Fprintf(os.Stderr, "  -m\tto count words\n")
		fmt.Fprintf(os.Stderr, "\nIf no flags are specified, defaults to counting lines, words, and bytes.\n")
		fmt.Fprintf(os.Stderr, "If no file is specified, reads from standard input.\n")
	}

	flag.BoolVar(&isByteMode, "c", false, "to count number of bytes")
	flag.BoolVar(&isLineMode, "l", false, "to count number of lines")
	flag.BoolVar(&isWordMode, "m", false, "to count number of words")

	flag.Parse()

	switch flag.NArg() {
	case 0:
		readingFromFile = false
		reader = os.Stdin
	case 1:
		readingFromFile = true
		fileName = flag.Arg(0)
	default:
		fmt.Fprintf(os.Stderr, "Error: too many arguments\n\n")
		flag.Usage()
		os.Exit(1)
	}
}

func main() {
	if readingFromFile {
		var err error
		reader, err = os.Open(fileName)
		checkForError(err)
	}

	switch {
	case isByteMode:
		fmt.Println(getNumberOf(reader, bufio.ScanBytes), fileName)
	case isLineMode:
		fmt.Println(getNumberOf(reader, bufio.ScanLines), fileName)
	case isWordMode:
		fmt.Println(getNumberOf(reader, bufio.ScanWords), fileName)
	default:
		buff, err := io.ReadAll(reader)
		checkForError(err)

		for _, mode := range []bufio.SplitFunc{bufio.ScanLines, bufio.ScanWords, bufio.ScanBytes} {
			fmt.Printf("%v ", getNumberOf(bytes.NewBuffer(buff), mode))
		}
		fmt.Printf("%s\n", fileName)
	}
}

func checkForError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		os.Exit(1)
	}
}

func getNumberOf(reader io.Reader, splitMode bufio.SplitFunc) int64 {
	var count int64
	scanner := bufio.NewScanner(reader)
	scanner.Split(splitMode)

	for scanner.Scan() {
		count++
	}
	checkForError(scanner.Err())

	return count
}
