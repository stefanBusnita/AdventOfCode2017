package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

func main() {

	f, err := os.Open("passwords.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	validPasswords := 0
	for scanner.Scan() {
		row := scanner.Text()
		fmt.Printf("Processing row: %s \n", row)
		words := strings.Fields(row)

		if checkValidPass(words, row) == len(words) {
			validPasswords++
		}
	}
	fmt.Printf("There are %d valid passowrds ", validPasswords)
}

// checkValidPass will check if the password is valid if there are no words that appear more than once
func checkValidPass(words []string, row string) int {
	wordCount := 0
	for i := 0; i < len(words); i++ {
		// check for the word exactly
		regEx, err := regexp.Compile(`\b` + words[i] + `\b`)
		check(err)
		// in the whole row
		found := regEx.FindAllStringIndex(row, -1)
		// each must be found once
		if len(found) == 1 {
			wordCount++
		}
	}
	return wordCount
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
