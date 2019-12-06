package hangman

import (
	"strings"
	"fmt"
)

type Game struct {
	State string
	Letters []string
	FoundLetters []string
	UsedLetters []string
	TurnsLeft int
}

func New(turns int, word string) (*Game, error)  {
	if len(word) < 3 {
		return nil, fmt.Errorf("Min word 3 char")
	}

	letters := strings.Split(strings.ToUpper(word), "")
	found := make([]string, len(letters))
	for i := 0; i < len(found); i++ {
		found[i] = "_"
	}

	return &Game{
		State:        "",
		Letters:      letters,
		FoundLetters: found,
		UsedLetters:  []string{},
		TurnsLeft:    turns,
	}, nil
}

func (g *Game) MakeAGuess(guess string)  {
	guess = strings.ToUpper(guess)

	if letterInWord(guess, g.UsedLetters) {
		g.State = "alreadyGuessed"
	} else if letterInWord(guess, g.Letters) {
		g.State = "goodGuess"
		g.RevealLetter(guess)

		if hasWon(g.Letters, g.FoundLetters) {
			g.State = "won"
		}
	} else {
		g.State = "badGuess"
		g.UsedLetters = append(g.UsedLetters, guess)
		g.TurnsLeft -= 1
		if g.TurnsLeft <= 0 {
			g.State = "lost"
		}
	}
}

func (g *Game) RevealLetter(guess string) {
	g.UsedLetters = append(g.UsedLetters, guess)
	for i, letter := range g.Letters {
		if letter == guess {
			g.FoundLetters[i] = letter
		}
	}
}

func letterInWord(guess string, letters []string) bool  {
	for _, l := range letters {
		if l == guess {
			return true
		}
	}

	return false
}

func hasWon(letters []string, foundLetter []string) bool  {
	for i := range letters {
		if letters[i] != foundLetter[i] {
			return false
		}
	}

	return true
}
