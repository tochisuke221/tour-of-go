package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func main() {
	showLineNumbers := flag.Bool("n", false, "show line numbers")
	flag.Parse()

	filePaths := flag.Args()

	if len(filePaths) == 0 {
		fmt.Println("Usage: mycat [-n] file1 file2 ...")
		
		return
	}

	lineNumber := 1

	for _, filePath := range filePaths {
		file, err := os.Open(filePath)
		
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error opening file %s: %v\n", filePath, err)
			continue
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if *showLineNumbers {
				fmt.Printf("%d: %s\n", lineNumber, line)
				lineNumber++
			} else {
				fmt.Println(line)
			}
		}

		// エラーチェック
		if err := scanner.Err(); err != nil {
			fmt.Fprintf(os.Stderr, "Error reading file %s: %v\n", filePath, err)
		}
		

	}



}