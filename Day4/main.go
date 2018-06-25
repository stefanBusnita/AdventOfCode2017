package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {

	f, err := os.Open("passwords.txt")
	check(err)
	defer f.Close()
	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	validPasswords := 0
	validAnagramPass := 0
	for scanner.Scan() {
		row := scanner.Text()
		fmt.Printf("Processing row: %s \n", row)
		words := strings.Fields(row)

		if checkValidPass(words, row) {
			validPasswords++
		}
		if checkForAnagrams(words, row) {
			validAnagramPass++
		}
	}
	fmt.Printf("There are %d valid passowrds ", validPasswords)
	fmt.Printf("There are %d valid anagram free passowrds ", validAnagramPass)
}

// checkValidPass will check if the password is valid if there are no words that appear more than once
func checkValidPass(words []string, row string) bool {
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
	return wordCount == len(words)
}

// checkForAnagrams will check if the password is valid if there are no words than have anagrams in the sentence (at most one for every row)
func checkForAnagrams(words []string, row string) bool {
	anagrams := 0
	for i := 0; i < len(words); i++ {
		currentWord := words[i]
		for j := 0; j < len(words); j++ {
			if areAnagrams(currentWord, words[j]) {
				anagrams++
			}
		}
	}
	return anagrams == len(words)
}

// Considering there are no whitespaces and UTF-8 characters we can use
// a nice workaround for checking anagrams with runes
type runes []rune

func (r runes) Len() int           { return len(r) }
func (r runes) Less(i, j int) bool { return r[i] < r[j] }
func (r runes) Swap(i, j int)      { r[i], r[j] = r[j], r[i] }

//  sortString will sort the string according to the above defined rules
func sortString(w string) string {
	r := []rune(w)
	sort.Sort(runes(r))
	return string(r)
}

// areAnagrams checks if the sorted strings are equal
func areAnagrams(word1, word2 string) bool {
	sortedWord1 := sortString(word1)
	sortedWord2 := sortString(word2)
	return sortedWord1 == sortedWord2
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}
