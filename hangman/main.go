package main

import (
	"fmt"
	"os"
	"training.go/hangman/hangman"
)

func main() {
	g := hangman.New(8,"test")

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
	}
}
