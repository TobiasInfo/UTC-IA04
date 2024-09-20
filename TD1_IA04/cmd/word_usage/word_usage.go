package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

// ########################################
// #                                      #
// #            PALINDROME                #
// #                                      #
// ########################################

// IsPalindrome returns true if the string s is a palindrome
func IsPalindrome(s string) bool {
	var n int = len(s)
	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}
	return true
}

// Palindromes returns a slice of palindromes found in the dictionary
func Palindromes(dict []string) []string {
	var palindromes []string
	for _, word := range dict {
		if IsPalindrome(word) {
			palindromes = append(palindromes, word)
		}
	}
	return palindromes
}


// ########################################
// #                                      #
// #            ANAGRAMS                  #
// #                                      #
// ########################################

// FootPrint returns the foot print of a string
func FootPrint(s string) string {
	// Split the string into characters
	var runes []rune = []rune(s)

	// Sort the characters
	sort.Slice(runes, func(i, j int) bool {
		return runes[i] < runes[j]
	})

	// Concatenate the characters
	var concat string = ""
	for _, c := range runes {
		concat += string(c)
	}

	return concat
}

func FootPrint2(s string) string {
	// Convert the string to a slice of strings, where each string is a single character
	var chars []string = strings.Split(s, "")

	// Sort the slice of strings
	sort.Strings(chars)

	// Join the sorted slice back into a single string
	return strings.Join(chars, "")
}

func Anagrams(words []string) (anagrams map[string][]string) {

	anagrams = make(map[string][]string)

	for _, word := range words {
		footprint := FootPrint(word)
		anagrams[footprint] = append(anagrams[footprint], word)
	}

	return
}


// ########################################
// #                                      #
// #            FILE I/O                  #
// #                                      #
// ########################################

func DictFromFile(filename string) (dict []string) {
	data, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println("Error reading file:", err)
		return
	}
	dict = strings.Split(string(data), "\n")
	return
}


// ########################################
// #                                      #
// #            MAIN                      #
// #                                      #
// ########################################

func main() {
	dict := [...]string{"AGENT", "CHIEN", "COLOC", "ETANG", "ELLE", "GEANT", "NICHE", "RADAR"}
	for _, word := range dict {
		fmt.Println(word, "is a palindrome:", IsPalindrome(word))
	}

	for _, word := range dict {
		if IsPalindrome(word) {
			fmt.Println("Palindrome found:", word)
		}
	}

	palindromes := Palindromes(dict[:])
	fmt.Println("Palindromes found:", palindromes)
	fmt.Println(FootPrint("AGENT"))
	fmt.Println(FootPrint2("AGENT"))
	fmt.Println(Anagrams(dict[:]))

	// var test []string = DictFromFile("dico-scrabble-fr.txt")
	// fmt.Println(test)
}
