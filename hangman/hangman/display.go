package hangman

import "fmt"

func DrawWelcome()  {
	fmt.Println(`
__ __   ____  ____    ____  ___ ___   ____  ____
|  |  | /    ||    \  /    ||   |   | /    ||    \
|  |  ||  o  ||  _  ||   __|| _   _ ||  o  ||  _  |
|  _  ||     ||  |  ||  |  ||  \_/  ||     ||  |  |
|  |  ||  _  ||  |  ||  |_ ||   |   ||  _  ||  |  |
|  |  ||  |  ||  |  ||     ||   |   ||  |  ||  |  |
|__|__||__|__||__|__||___,_||___|___||__|__||__|__|`)
}

func Draw(g *Game, guess string)  {
	drawTurns(g.TurnsLeft)
	drawState(g, guess)
}

func drawTurns(l int) {
	switch l {
	case 0:
		fmt.Println("Perdu connard")
	default:
		fmt.Printf("Plus que %d coups\n", l)
	}
}

func drawState(g *Game, guess string)  {
	fmt.Print("Guessed: ")
	drawLetters(g.FoundLetters)

	fmt.Print("Guessed: ")
	drawLetters(g.UsedLetters)

	switch g.State {
	case "goodGuess":
		fmt.Print("Good guess !")
	case "alreadyGuessed":
		fmt.Printf("Letter '%s' is already used", guess)
	case "badGuess":
		fmt.Printf("Bad guess '%s' is not int he world", guess)
	case "lost":
		fmt.Print("You lost bitch, the world was: ")
		drawLetters(g.Letters)
	case "won":
		fmt.Print("You won, the world was: ")
		drawLetters(g.Letters)
	}
}

func drawLetters(l []string)  {
	for _, c := range l {
		fmt.Printf("%v ", c)
	}
	fmt.Println()
}
