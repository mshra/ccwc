package main

import (
	"bufio"
	"bytes"
	"flag"
	"fmt"
	"io"
	"log"
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
	flag.BoolVar(&isByteMode, "c", false, "to read number of bytes")
	flag.BoolVar(&isLineMode, "l", false, "to read number of lines")
	flag.BoolVar(&isWordMode, "m", false, "to read number of words")

	flag.Parse()
	args := flag.Args()

	readingFromFile = false
	reader = os.Stdin

	if len(args) > 0 {
		fileName = args[0]
		readingFromFile = true
	}
}

func checkForError(err error) {
	if err != nil {
		log.Fatal(err)
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
