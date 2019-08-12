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
var incorrectly_guessed int
var word string

func empty_slice(guess []byte) {
	for i, _ := range guess {
		guess[i] = '_'
	}
}

func print_hangman(guess []byte) {
	fmt.Println(hangman_art[incorrectly_guessed])
	fmt.Printf(">> ")
	for _, c := range guess {
		fmt.Printf("%c ", c)
	}
	fmt.Printf("\n")
}

func update_guess(guess []byte, letter string){
	// Check if letter has been guessed before.
	index := strings.Index(word, letter)
	if index != -1 {
		for i := index; i < len(word); i++ {
			// If letter has already been guessed.
			if word[i] == guess[i] {
				break
			}
			if word[i] == letter[0] {
				guess[i] = letter[0]
				correctly_guessed++
			}
		}
	} else {
		incorrectly_guessed++
		fmt.Println("Letter not found.")
	}
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
	word = words[rand_index]
	// word = words[0]
	word = word[:len(word) - 1]

	// Setting up the game.
	correctly_guessed = 0
	incorrectly_guessed = 0

	guess_slice := []byte(word)
	empty_slice(guess_slice)

	for {
		print_hangman(guess_slice)

		if correctly_guessed == len(word) {
			break
		}

		if incorrectly_guessed == 5 {
			fmt.Printf(">> Boohoo, game over!")
			break
		}

		fmt.Println(">> Enter a letter.")
		letter, _ := reader.ReadString('\n')

		// Checking if input is valid.
		if len(letter) > 3 {
			continue
		} 

		// Check if the letter exists in the word now
		update_guess(guess_slice, letter[:1])
	}
}