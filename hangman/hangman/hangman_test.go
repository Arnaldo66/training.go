package hangman

import "testing"


func TestLetterInWord(t *testing.T)  {
    word := []string{"b", "o", "b"}
    guess := "b"

    hasLetter := letterInWord(guess, word)
    if !hasLetter {
        t.Errorf("Not ok")
    }
}

func TestLetterNotInWord(t *testing.T)  {
    word := []string{"b", "o", "b"}
    guess := "r"

    hasLetter := letterInWord(guess, word)
    if hasLetter {
        t.Errorf("Not ok")
    }
}

func TestInvalidWord(t *testing.T)  {
    _, err := New(1, "")
    if err == nil {
        t.Errorf("Not ok")
    }
}

func TestGameGoodGuess(t *testing.T)  {
    g, _ := New(3, "bob")
    g.MakeAGuess("b")
    validateStates(t, "goodGuess", g.State)
}

func TestGameAlreadyGuess(t *testing.T)  {
    g, _ := New(3, "bob")
    g.MakeAGuess("b")
    g.MakeAGuess("b")
    validateStates(t, "alreadyGuessed", g.State)
}

func TestGameBadGuess(t *testing.T)  {
    g, _ := New(3, "bob")
    g.MakeAGuess("u")
    validateStates(t, "badGuess", g.State)
}

func TestGameWonGuess(t *testing.T)  {
    g, _ := New(3, "bob")
    g.MakeAGuess("b")
    g.MakeAGuess("o")
    validateStates(t, "won", g.State)
}

func TestGameLostGuess(t *testing.T)  {
    g, _ := New(1, "bob")
    g.MakeAGuess("u")
    validateStates(t, "lost", g.State)
}

func validateStates(t *testing.T, expectedState, currentState string)  bool {
    if expectedState != currentState {
        t.Errorf("State are not ok")
        return false
    }

    return true
}
