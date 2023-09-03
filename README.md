# Golang Spellcheck
This is a spell checker for the Go programming language. When I looked for a spell checker for Go, I couldn't find one, so I started this repository. We can all work on this to make it better. I've created the initial basic version. When I used the code here to spell check a document I had already run through Google Drive and Grammarly, I found a legitimate spelling error that both had missed (personaes should be personae), so good progress has been made!

# English words and names
**I found great sources for all of the words in English.** From Google "The Second Edition of the 20-volume Oxford English Dictionary contains full entries for 171,476 words in current use (and 47,156 obsolete words). Webster's Third New International Dictionary, Unabridged, together with its 1993 Addenda Section, includes some 470,000 entries." From the sources I've found, listed below, there are currently 372,933 unique words and names. For the names, I used the social security administration's records of names given at birth. **You can see the sources of the words and names below.**

# Next steps

1. Improve spell checker
	i. refactor, packages, more robust code, etc
2. Create a grammar checker

This is all from chatGPT - for inspiration:

### How would you write a spell checker in the go programming language?
A spell checker can be a complex system depending on how advanced you want it to be. The simplest form of a spell checker can just look up a given word in a dictionary to check if it is correctly spelled. More advanced spell checkers also consider contextual information to differentiate between words that are spelled correctly but used incorrectly.

Here's a basic example of a spell checker in Go, using a set data structure to hold a dictionary of correct words:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Initialize a dictionary set
var dictionary = map[string]struct{}{
	"apple":   {},
	"orange":  {},
	"banana":  {},
	"pear":    {},
	"pineapple": {},
}

// Function to check if a word is correctly spelled
func isCorrectlySpelled(word string) bool {
	_, exists := dictionary[word]
	return exists
}

func main() {
	// Read text from user
	fmt.Println("Enter some text:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))

	// Split text into words
	words := strings.Fields(text)

	// Check each word for spelling
	for _, word := range words {
		word = strings.Trim(word, ".,!?\"'()[]{}:;")

		if isCorrectlySpelled(word) {
			fmt.Printf("'%s' is correctly spelled.\n", word)
		} else {
			fmt.Printf("'%s' is misspelled.\n", word)
		}
	}
}
```

To run this example:

1. Save the code into a file, such as `spellchecker.go`.
2. Run `go run spellchecker.go`.
3. Enter some text to check the spelling of each word.

This example only uses a very small hardcoded dictionary. In a real-world application, you would want to populate your dictionary from a more comprehensive source. Also, this example doesn't handle many complexities like plural forms, past tense, etc. Advanced spell checkers often use machine learning models or advanced algorithms for such purposes.

You can extend this basic example to:

- Add more words to your dictionary.
- Handle different forms of words (e.g., stemming).
- Use algorithms like Levenshtein distance to suggest corrections.
- Use a pre-built package to improve accuracy.
- Include contextual information to check for the correct usage of words.

For instance, you could use the Levenshtein distance algorithm to suggest the closest words from the dictionary for a misspelled word. But for the sake of simplicity, the above example shows the basic spell checking mechanism.

### Can you add algorithms like Levenshtein distance to suggest corrections to the program?

Certainly! Below is an extended version of the spell checker in Go that uses the Levenshtein distance algorithm to suggest the closest word from the dictionary for a misspelled word.

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Initialize a dictionary set
var dictionary = map[string]struct{}{
	"apple":    {},
	"orange":   {},
	"banana":   {},
	"pear":     {},
	"pineapple": {},
}

// Function to check if a word is correctly spelled
func isCorrectlySpelled(word string) bool {
	_, exists := dictionary[word]
	return exists
}

// Calculate the Levenshtein distance between two strings
func levenshtein(a, b string) int {
	la, lb := len(a), len(b)
	d := make([][]int, la+1)
	for i := 0; i <= la; i++ {
		d[i] = make([]int, lb+1)
		d[i][0] = i
	}
	for j := 0; j <= lb; j++ {
		d[0][j] = j
	}

	for i := 1; i <= la; i++ {
		for j := 1; j <= lb; j++ {
			cost := 0
			if a[i-1] != b[j-1] {
				cost = 1
			}
			d[i][j] = min(d[i-1][j]+1, d[i][j-1]+1, d[i-1][j-1]+cost)
		}
	}
	return d[la][lb]
}

// Min function for integers
func min(nums ...int) int {
	m := nums[0]
	for _, n := range nums[1:] {
		if n < m {
			m = n
		}
	}
	return m
}

// Find the closest word from the dictionary
func findClosest(word string) string {
	closestWord := ""
	minDistance := 1<<31 - 1 // Initialize to max int value

	for dictWord := range dictionary {
		distance := levenshtein(word, dictWord)
		if distance < minDistance {
			closestWord = dictWord
			minDistance = distance
		}
	}

	return closestWord
}

func main() {
	// Read text from user
	fmt.Println("Enter some text:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))

	// Split text into words
	words := strings.Fields(text)

	// Check each word for spelling
	for _, word := range words {
		word = strings.Trim(word, ".,!?\"'()[]{}:;")

		if isCorrectlySpelled(word) {
			fmt.Printf("'%s' is correctly spelled.\n", word)
		} else {
			fmt.Printf("'%s' is misspelled. Did you mean '%s'?\n", word, findClosest(word))
		}
	}
}
```

This implementation utilizes the Levenshtein distance algorithm to find the closest word in the dictionary for each misspelled word. The function `levenshtein` computes the distance between two strings. The `findClosest` function uses this algorithm to find the closest word from the dictionary for a given word.

To run this example:

1. Save the code into a file, such as `spellchecker_with_suggestions.go`.
2. Run `go run spellchecker_with_suggestions.go`.
3. Enter some text to check the spelling of each word.

The program will now suggest the closest word from the dictionary when a misspelled word is encountered. Keep in mind that this is a simple example and may not cover all cases or be the most efficient implementation.

### Starting from the first code example you provided, and not including the levenshetein computation, can you add contextual information to check for the correct usage of words?

Certainly! Adding contextual information can make your spell checker much smarter in identifying words that are spelled correctly but used incorrectly in a sentence. One common way to do this is to use n-grams. An n-gram is a contiguous sequence of n items from a given sample of text. In this example, I'll use bigrams (n=2) to identify the incorrect usage of words based on their immediate neighbor.

Firstly, let's create a simple bigram model. For demonstration, I'm using hard-coded bigrams, but in a real application, you'd train this on a large dataset.

Here's a version of the initial spell checker with basic contextual checking:

```go
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// Initialize a dictionary set
var dictionary = map[string]struct{}{
	"apple":   {},
	"eat":     {},
	"orange":  {},
	"banana":  {},
	"pear":    {},
	"pineapple": {},
	"i":       {},
	"love":    {},
}

// Initialize bigram model for contextual spelling
var bigrams = map[string][]string{
	"i":       {"love", "eat"},
	"love":    {"apple", "banana", "pear", "pineapple", "orange"},
	"eat":     {"apple", "banana", "pear", "pineapple", "orange"},
	"apple":   {},
	"banana":  {},
	"pear":    {},
	"pineapple": {},
	"orange":  {},
}

// Function to check if a word is correctly spelled
func isCorrectlySpelled(word string) bool {
	_, exists := dictionary[word]
	return exists
}

// Function to check if a bigram is valid according to the bigram model
func isValidBigram(word1, word2 string) bool {
	nextWords, exists := bigrams[word1]
	if !exists {
		return false
	}

	for _, nextWord := range nextWords {
		if nextWord == word2 {
			return true
		}
	}

	return false
}

func main() {
	// Read text from user
	fmt.Println("Enter some text:")
	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.ToLower(strings.TrimSpace(text))

	// Split text into words
	words := strings.Fields(text)

	// Check each word for spelling
	for i, word := range words {
		word = strings.Trim(word, ".,!?\"'()[]{}:;")

		if isCorrectlySpelled(word) {
			fmt.Printf("'%s' is correctly spelled.\n", word)

			// Check if the word is used correctly in context
			if i > 0 {
				prevWord := strings.Trim(words[i-1], ".,!?\"'()[]{}:;")
				if !isValidBigram(prevWord, word) {
					fmt.Printf("However, '%s %s' seems like an incorrect usage.\n", prevWord, word)
				}
			}
		} else {
			fmt.Printf("'%s' is misspelled.\n", word)
		}
	}
}
```

In this example, I've created a simple bigram model (`bigrams`) to map each word to a slice of words that can validly follow it. The function `isValidBigram` checks if a pair of words form a valid bigram according to the model.

To run this example:

1. Save the code into a file, such as `contextual_spellchecker.go`.
2. Run `go run contextual_spellchecker.go`.
3. Enter some text to check the spelling and contextual usage of each word.

This example uses very basic contextual checking and can definitely be extended further. For instance, you could use more advanced NLP techniques or machine learning models to do more accurate contextual spelling checks.

# English words and names sources
--------------------------
files

sowpods.txt
https://www.wordgamedictionary.com/sowpods/

enable.txt
https://www.wordgamedictionary.com/enable/

twl06.txt
https://www.wordgamedictionary.com/twl06/

mit-10000-most-common.txt
https://www.mit.edu/~ecprice/wordlist.10000

3000-most-common.txt
https://www.ef.com/wwen/english-resources/english-vocabulary/top-3000-words/

enable.txt
https://www.wordgamedictionary.com/dictionary/

big.txt
https://norvig.com/spell-correct.html

yob*.txt
https://www.ssa.gov/oact/babynames/limits.html

--------------------------
files-others

wordlist.txt
umich professor 1999
https://www-personal.umich.edu/~jlawler/wordlist.html

words.txt
https://github.com/dwyl/english-words/blob/master/words.txt

wordlist-english.txt
https://gist.github.com/MarvinJWendt/a7fb66c187adceaa049eab49f1815260

words2.txt
https://gist.github.com/fay59/53f885f1fc5856741cb4

gists
https://gist.github.com/search?q=english+words
