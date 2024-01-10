package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
)

const (
	
	asciiArt = `
                        __
   _________  _  ______/ /
  / ___/ __ \/ |/_/ __  / 
 / /  / /_/ />  </ /_/ /  
/_/   \__, /_/|_|\__,_/   
     /____/               
`
		)

func applyRegex(input string, regexList []*regexp.Regexp) string {
	result := input
	for _, regex := range regexList {
		result = regex.ReplaceAllString(result, "")
	}
	return result
}

func main() {
	
	fmt.Println(asciiArt)
	fmt.Println("\nFormat: [Input payload] -> [Processed payload]\n")	

	if len(os.Args) != 3 {
		fmt.Println("Usage: rgxd payloads regex")
		os.Exit(1)
	}

	inputFileName := os.Args[1]
	regexFileName := os.Args[2]

	inputFile, err := os.Open(inputFileName)
	if err != nil {
		fmt.Println("Error opening input file:", err)
		os.Exit(1)
	}
	defer inputFile.Close()

	regexFile, err := os.Open(regexFileName)
	if err != nil {
		fmt.Println("Error opening regex file:", err)
		os.Exit(1)
	}
	defer regexFile.Close()

	var regexList []*regexp.Regexp
	scanner := bufio.NewScanner(regexFile)
	for scanner.Scan() {
		regexPattern := scanner.Text()
		compiledRegex, err := regexp.Compile(regexPattern)
		if err != nil {
			fmt.Println("Error compiling regex:", err)
			os.Exit(1)
		}
		regexList = append(regexList, compiledRegex)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading regex file:", scanner.Err())
		os.Exit(1)
	}

	scanner = bufio.NewScanner(inputFile)
	for scanner.Scan() {
		inputString := scanner.Text()
		outputString := applyRegex(inputString, regexList)
		fmt.Printf("[%s] -> [%s]\n", inputString, outputString)
	}

	if scanner.Err() != nil {
		fmt.Println("Error reading input file:", scanner.Err())
		os.Exit(1)
	}
}

