package hangman

import "strings"

type Game struct {
	State string
	Letters []string
	FoundLetters []string
	UsedLetters []string
	TurnsLeft int
}

func New(turns int, word string) *Game  {
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
	}
}