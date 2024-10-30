package main

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
)

var isNumberOfByteModeEnabled bool
var isNumberOfLineModeEnabled bool
var fileName string

func init() {
	flag.BoolVar(&isNumberOfByteModeEnabled, "c", false, "read number of bytes in the given file")
	flag.BoolVar(&isNumberOfLineModeEnabled, "l", false, "read number of lines in the given file")

	flag.Parse()
}

func main() {
	fileName = os.Args[len(os.Args)-1]
	file, err := os.Open(fileName)
	checkError(err)
	defer file.Close()

	if isNumberOfByteModeEnabled {
		fmt.Println(getFileStat(file).Size(), fileName)
	}

	if isNumberOfByteModeEnabled {
		fmt.Println(getNumberOfLinesInFile(file), fileName)
	}
}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getFileStat(f *os.File) os.FileInfo {
	fileStat, err := f.Stat()
	checkError(err)
	return fileStat
}

func getNumberOfLinesInFile(f *os.File) int {
	var lineCount int
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		lineCount++
	}

	return lineCount
}
