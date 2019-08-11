package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"io/ioutil"
	"math/rand"
)

var correctly_guessed int

func empty_slice(guess []byte) {
	for i := 0; i < len(guess); i++ {
		guess[i] = '_'
	}
}

func print_hangman(guess []byte) {
	for i := 0; i < len(guess); i++ {
		fmt.Printf("%c ", guess[i])
	}
	fmt.Printf("\n")
}

func update_guess(guess []byte, letter string, word string) []byte {
	fmt.Println(guess)
	fmt.Println(letter)
	fmt.Println(word)
	// Check if letter has been guessed before.

	if strings.Contains(word, letter) {
		for i := 0; i <= len(word); i++ {
			if word[i] == letter[0] {
				guess[i] = letter[0]
				correctly_guessed++
			}
		}
	}

	return guess
}

func main() {
	fmt.Println(" ")
	fmt.Println(">> Welcome to Hangman in Go!")
	fmt.Println(">> Are you ready to play (Y/N) ?")

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')

	// Is there an easier way to remove the pesky CR at the end?
	if strings.Compare("N", text[:1]) == 0 {
		fmt.Println(">> Why did you start me up if you didn't want to play, gee?")
		fmt.Println(">> Bye bye now.")
		return
	}

	fmt.Println(">> Great, let's begin!")

	// Reading in list of words.
	content, err := ioutil.ReadFile("hangman-words.txt")
	if err != nil {
		fmt.Println(">> Uh oh, error reading in the words file!")
		return
	}
	words := strings.Split(string(content), "\n")

	// Selecting a word.
	rand_index := rand.Intn(len(words))
	word := words[rand_index]
	fmt.Println("The word is: ", word)

	correctly_guessed = 0
	guess_slice := []byte(word)
	empty_slice(guess_slice)
	return

	for {
		fmt.Println("The word is: ")
		print_hangman(guess_slice)

		if correctly_guessed == len(word) {
			break
		}

		fmt.Println(">> Enter a letter.")
		letter, _ := reader.ReadString('\n')

		// Checking if input is valid.
		if len(letter) > 2 {
			// fmt.Println(">> Enter only one letter. Try again.")
			continue
		} 

		// Check if the letter exists in the word now
		guess_slice = update_guess(guess_slice, letter[:1], word)
	}
}