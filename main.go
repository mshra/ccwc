package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var (
	fileName                   string
	readingFromStdin           bool
	isNumberOfByteModeEnabled  bool
	isNumberOfLineModeEnabled  bool
	isNumberOfWordsModeEnabled bool
)

func init() {
	flag.BoolVar(&isNumberOfByteModeEnabled, "c", false, "read number of fileToReadNumberOfBytes in the given file")
	flag.BoolVar(&isNumberOfLineModeEnabled, "l", false, "read number of fileToReadNumberOfLines in the given file")
	flag.BoolVar(&isNumberOfWordsModeEnabled, "m", false, "read number of fileToReadNumberofWords in the given file")

	flag.Parse()
	args := flag.Args()

	readingFromStdin = true
	if len(args) > 0 {
		fileName = args[0]
		readingFromStdin = false
	}
}

func main() {
	if isNumberOfByteModeEnabled {
		file, err := os.Open(fileName)
		checkError(err)
		defer file.Close()
		fmt.Println(getNumberOfBytesInFile(file), fileName)
		return
	}

	if isNumberOfLineModeEnabled {
		file, err := os.Open(fileName)
		checkError(err)
		defer file.Close()
		fmt.Println(getNumberOfLinesInFile(file), fileName)
		return
	}

	if isNumberOfWordsModeEnabled {
		file, err := os.Open(fileName)
		checkError(err)
		defer file.Close()
		fmt.Println(getNumberOfWordsInFile(file), fileName)
		return
	}

	fileToReadNumberOfBytes, err := os.Open(fileName)
	checkError(err)
	defer fileToReadNumberOfBytes.Close()

	fileToReadNumberOfLines, err := os.Open(fileName)
	checkError(err)
	defer fileToReadNumberOfLines.Close()

	fileToReadNumberofWords, err := os.Open(fileName)
	checkError(err)
	defer fileToReadNumberofWords.Close()

	fmt.Println(getNumberOfLinesInFile(fileToReadNumberOfLines), getNumberOfWordsInFile(fileToReadNumberofWords), getNumberOfBytesInFile(fileToReadNumberOfBytes), fileName)
	return
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getNumberOfBytesInFile(f *os.File) int64 {
	fileStat, err := f.Stat()
	checkError(err)
	return fileStat.Size()
}

func getNumberOfLinesInFile(f *os.File) int64 {
	var lineCount int64
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}

func getNumberOfWordsInFile(f *os.File) int {
	wordsCount := 0
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordsCount++
	}
	checkError(scanner.Err())

	return wordsCount
}
