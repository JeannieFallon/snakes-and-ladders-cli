/* Implement board game known as Chutes & Ladders,
based on Indian game Snakes & Ladders:
https://en.wikipedia.org/wiki/Snakes_and_Ladders

Currently single-player mode only with 5x5 board. */

package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {

	// initialize starting values
	currSpace := 0
	rollVal := 0

	// initialize unique game board & print
	board := getBoard()
	printBoard(board, currSpace)

	scanner := bufio.NewScanner(os.Stdin)
	for currSpace < 25 {
		// roll die
		fmt.Println("Press enter to roll")
		scanner.Scan()
		rollVal = rollDie()
		fmt.Printf("You rolled: %d\n", rollVal)

		// apply roll if valid, and apply board condition if present
		if checkForValidMove(rollVal, currSpace) {
			currSpace = applyRoll(rollVal, currSpace)
			currSpace = applyBoardCondition(board, currSpace)
		} else {
			fmt.Printf("Invalid move. Remain on current space and roll again.\n")
		}
		printBoard(board, currSpace)
	}

	fmt.Println("End Chutes")
}

/* generate & display unique game board */

func getBoard() map[int]int {
	// generate unique board of 25 spaces total where 15 spaces have secondary move
	var key int
	var board map[int]int
	board = make(map[int]int)

	for i := 0; i < 15; i++ {
		// generate unique keys as 15 random numbers from 1 to 25
		for {
			key = getRandNum(25)
			if val, ok := board[key]; !ok {
				//TODO fix unused val
				val++
				break
			}
		}

		// set each map value where -5 <= n <=5
		//TODO exclude 0
		board[key] = getRandNum(10) - 5
	}

	return board
}

func printBoard(board map[int]int, currSpace int) {
	fmt.Printf("Current game board:\n")

	// decrementing loop to print board indices in reverse
	for i := 25; i > 0; i-- {
		if i%5 == 0 {
			fmt.Printf("\n")
		}

		// 	board legend:
		// 		X	current space of player's piece
		// 		+	space with secondary move forward
		// 		-	space with secondary move back
		// 		. 	space with no secondary move

		if i == currSpace {
			fmt.Printf(" X ")
		} else {
			val := board[i]
			if val > 0 {
				fmt.Printf(" + ")
			} else if val < 0 {
				fmt.Printf(" - ")
			} else {
				fmt.Printf(" . ")
			}
		}
	}

	//TODO fix redundant newline formatting bs
	fmt.Printf("\n\n")
}

/* generate & process game move */

func rollDie() int {
	//TODO test & fix range if needed (don't include 0)
	// hardcode for 6-sided die
	return getRandNum(7)
}

func applyRoll(rollVal, currSpace int) int {
	return currSpace + rollVal
}

func applyBoardCondition(board map[int]int, currSpace int) int {
	// check if current space exists as key and add val to current space if exists
	key := currSpace
	if currSpace, ok := board[currSpace]; ok {
		var directionMsg string
		var directionEval bool
		secondaryMove := board[currSpace]
		// only valid if enough spaces left on board for entire move
		if checkForValidMove(secondaryMove, currSpace) {
			currSpace += secondaryMove
			// reset any negative decrement to starting space
			if currSpace < 0 {
				currSpace = 0
			}
			// set messaging
			switch directionEval {
			case secondaryMove < 0:
				directionMsg = "BACK"
			case secondaryMove == 0:
				directionMsg = "NO"
			case secondaryMove > 0:
				directionMsg = "FORWARD"
			}
			fmt.Printf("You landed on space: %d\n", currSpace)
			fmt.Printf("This space requires %s move by %d spaces\n", directionMsg, abs(board[key]))
		} else {
			fmt.Printf("Secondary move is invalid. Remain on current space.\n")
		}
	} else {
		fmt.Printf("No secondary move. Remain on current space.\n")
	}

	return currSpace

}

func checkForValidMove(rollVal int, currSpace int) bool {
	return rollVal+currSpace < 25
}

/* helpers */

func getRandNum(upperBound int) int {
	rand.Seed(time.Now().UnixNano())
	return rand.Intn(upperBound)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
