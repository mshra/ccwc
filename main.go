package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"log"
	"os"
)

var (
	fileName         string
	readingFromStdin bool
	isByteMode       bool
	isLineMode       bool
	isWordMode       bool
)

func init() {
	flag.BoolVar(&isByteMode, "c", false, "to read number of bytes")
	flag.BoolVar(&isLineMode, "l", false, "to read number of lines")
	flag.BoolVar(&isWordMode, "m", false, "to read number of words")

	flag.Parse()
	args := flag.Args()

	readingFromStdin = true
	if len(args) > 0 {
		fileName = args[0]
		readingFromStdin = false
	}
}

func main() {
	file, err := os.Open(fileName)
	checkForError(err)
	defer file.Close()

	if isByteMode {
		numberOfBytes := getNumberOf(file, bufio.ScanBytes)
		fmt.Println(numberOfBytes, fileName)
		return
	}

	if isLineMode {
		numberOfLines := getNumberOf(file, bufio.ScanLines)
		fmt.Println(numberOfLines, fileName)
		return
	}

	if isWordMode {
		numberOfWords := getNumberOf(file, bufio.ScanWords)
		fmt.Println(numberOfWords, fileName)
		return
	}

	// fmt.Println(getNumberOfLinesInFile(file), getNumberOfWordsInFile(file), getNumberOfBytesInFile(file), fileName)
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
