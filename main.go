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
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	if isNumberOfByteModeEnabled {
		fmt.Println(getNumberOfBytesInFile(file), fileName)
		return
	}

	if isNumberOfLineModeEnabled {
		fmt.Println(getNumberOfLinesInFile(file), fileName)
		return
	}

	if isNumberOfWordsModeEnabled {
		fmt.Println(getNumberOfWordsInFile(file), fileName)
	}

	file.Seek(0, 0)
	fmt.Println(getNumberOfLinesInFile(file), getNumberOfWordsInFile(file), getNumberOfBytesInFile(file), fileName)
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

func getNumberOfWordsInFile(f *os.File) int64 {
	var wordsCount int64
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordsCount++
	}
	checkError(scanner.Err())

	return wordsCount
}

func testingNumberOfLinesFromStdin() int64 {
	var lines int64
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		lines += 1
	}

	return lines
}

func testigNumberofbytesinfile() int64 {
	var byteCount int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanBytes)

	for scanner.Scan() {
		byteCount++
	}
	checkError(scanner.Err())

	return byteCount
}

func testingNumberOfWordsInFile() int64 {
	var wordcount int64
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		wordcount++
	}
	checkError(scanner.Err())

	return wordcount
}
