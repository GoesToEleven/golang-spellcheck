package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"unicode"
)

const (
	ValidWordsFileName = "all-words.txt"
	FileToCheck        = "check-spelling.txt"
)

func main() {
	// Read all valid words
	allWords, err := readWords(ValidWordsFileName)
	if err != nil {
		log.Fatalf("Error reading %s: %v", ValidWordsFileName, err)
	}

	// Open the file to check
	fileToCheck, err := os.Open(FileToCheck)
	if err != nil {
		log.Fatalf("Error opening %s: %v", FileToCheck, err)
	}

	// Check if the file is empty
	fileInfo, err := fileToCheck.Stat()
	if err != nil {
		log.Fatalf("Error getting file information: %v", err)
	}
	if fileInfo.Size() == 0 {
		log.Println("File is empty. Nothing to check.")
		return
	}

	defer fileToCheck.Close()

	checkSpelling(fileToCheck, allWords)
}

// readWords reads a list of words from a file and returns a set.
func readWords(filename string) (map[string]bool, error) {
	file, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	words := make(map[string]bool)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		word := strings.TrimSpace(scanner.Text())
		words[word] = true
	}
	return words, scanner.Err()
}

func checkSpelling(f *os.File, allWords map[string]bool) {
	// Create a scanner and read the file line by line.
	scanner := bufio.NewScanner(f)
	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := strings.TrimSpace(scanner.Text())
		wordsInLine := strings.Fields(line)

		for _, word := range wordsInLine {
			word = cleanWord(word)

			_, exists := allWords[word]
			if word != "" && !exists {
				fmt.Printf("Line %d: Spelling error in word %s\n", lineNumber, word)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		log.Printf("Error reading %s: %v", FileToCheck, err)
	}
}

// cleanWord removes digits and specified special characters from the word.
func cleanWord(word string) string {

	word = strings.ToLower(word)
	// Remove leading and trailing undesired characters
	word = strings.Trim(word, "*.,!?\"'()[]{}:;#&+-/=$%<>@_|~")
	// Remove leading and trailing spaces
	word = strings.TrimSpace(word)

	// Discard word with undesired characters
	xs := []string{"\a", "\b", "\f", "\n", "\r", "\t", "\v", "\\", "\"", "-", ".", "=", "'", "`", "?", "$", "’", "“", "‘", "–"}
	for _, v := range xs {
		if strings.Contains(word, v) {
			return ""
		}
	}

	// Discard word with numeric digits
	for _, r := range word {
		if unicode.IsDigit(r) {
			return ""
		}
	}

	return word
}
