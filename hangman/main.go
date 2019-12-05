package main

import (
	"fmt"
	"os"
	"training.go/hangman/hangman"
	"training.go/hangman/dictionary"
)

func main() {
	err := dictionary.Load("words.txt")
	if err != nil {
		fmt.Printf("Cannot open file %v" err)
		os.Exit(1)
	}

	g := hangman.New(8,dictionary.PickWord())

	hangman.DrawWelcome()

	guess := ""
	for {
		hangman.Draw(g, guess)
		if g.State == "won" || g.State == "lost" {
			os.Exit(0)
		}

		l, err := hangman.ReadGuess()
		if err != nil {
			fmt.Printf("Could not read %v", err)
			os.Exit(1)
		}

		guess = l
		g.MakeAGuess(guess)
	}
}
