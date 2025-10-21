// https://workat.tech/machine-coding/practice/snake-and-ladder-problem-zgtac9lxwntg

package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"strings"
	"time"
)

type dice struct {
	minValue int
	maxValue int
	value    int
}

func (d *dice) roll() int {
	d.value = rand.Intn(d.maxValue-d.minValue+1) + d.minValue
	return d.value
}

type snake struct {
	head int
	tail int
}

func addNewSnake(head, tail int) *snake {
	return &snake{
		head: head,
		tail: tail,
	}
}

type ladder struct {
	start int
	end   int
}

func addNewLadder(start, end int) *ladder {
	return &ladder{
		start: start,
		end:   end,
	}
}

type board struct {
	snakes  []*snake
	ladders []*ladder
}

func addNewBoard(cmds []string) *board {
	board := &board{}

	snakeCount, _ := strconv.Atoi(cmds[0])
	var snakes []*snake
	for i := 1; i < snakeCount+1; i++ {
		positionArr := strings.Fields(cmds[i])
		head, _ := strconv.Atoi(positionArr[0])
		tail, _ := strconv.Atoi(positionArr[1])
		snakes = append(snakes, addNewSnake(head, tail))
	}
	board.snakes = snakes
	fmt.Println("adding snakes in the board. SnakeCount:", snakeCount)

	ladderCount, _ := strconv.Atoi(cmds[snakeCount+1])
	var ladders []*ladder
	for i := snakeCount + 2; i < snakeCount+ladderCount+2; i++ {
		positionArr := strings.Fields(cmds[i])
		start, _ := strconv.Atoi(positionArr[0])
		end, _ := strconv.Atoi(positionArr[1])
		ladders = append(ladders, addNewLadder(start, end))
	}
	board.ladders = ladders
	fmt.Println("adding ladders in the board. LadderCount:", ladderCount)

	return board
}

func (d *board) getAdjustedPosition(position int, player *player) (bool, int) {
	// check for snake
	var isAdjusted bool
	for _, snake := range d.snakes {
		if snake.head == position {
			fmt.Println("[OOPS] A snake bit", player.name, "at", snake.head, "and now", player.name, "is at", snake.tail)
			position = snake.tail
			isAdjusted = true
		}
	}
	for _, ladder := range d.ladders {
		if ladder.start == position {
			fmt.Println("[YAY]", player.name, "got a ladder at", ladder.start, "and now", player.name, "is at", ladder.end)
			position = ladder.end
			isAdjusted = true
		}
	}
	return isAdjusted, position
}

type player struct {
	name     string
	position int
	dice     *dice
	isWin    bool
}

func addNewPlayer(name string) *player {
	return &player{
		name:     name,
		position: 0,
		dice: &dice{
			minValue: 1,
			maxValue: 6,
		},
	}
}

//<player_name> rolled a <dice_value> and moved from <initial_position> to <final_position>

func (d *player) rollDice() int {
	return d.position + d.dice.roll()
}

func (d *player) movePiece(position int, everAdjusted bool) {
	if !everAdjusted {
		fmt.Println(d.name, "rolled a", d.dice.value, "and moved from", d.position, "to", position)
	}
	d.position = position
}

func (d *player) isPositionValid(position int) bool {
	if position > 100 {
		return false
	}
	return true
}

func (d *player) isWon() bool {
	if d.position == 100 {
		fmt.Println(d.name, "wins the game")
		return true
	}
	return false
}

type game struct {
	board   *board
	players []*player
}

func (d *game) init(cmds []string) {
	fmt.Println("Initializing the board")
	d.board = addNewBoard(cmds)

	snakeCount, _ := strconv.Atoi(cmds[0])
	ladderCount, _ := strconv.Atoi(cmds[snakeCount+1])
	playerCount, _ := strconv.Atoi(cmds[snakeCount+ladderCount+2])
	var players []*player
	for i := snakeCount + ladderCount + 3; i < snakeCount+playerCount+ladderCount+3; i++ {
		players = append(players, addNewPlayer(cmds[i]))
	}
	d.players = players
	fmt.Println("adding players to the game. PlayerCount:", playerCount)
}

func (d *game) start() {
	currPlayer := d.players[0]
	for !currPlayer.isWon() {
		currPlayer = d.players[0]
		// fmt.Println("This is", currPlayer.name, "turn")
		currPosition := currPlayer.rollDice()
		if !currPlayer.isPositionValid(currPosition) {
			fmt.Println("This is an invalid position,", currPlayer.name, "needs to play again")
			continue
		}

		var everAdjusted bool
		for {
			isAdjusted, adjustedPosition := d.board.getAdjustedPosition(currPosition, currPlayer)
			if isAdjusted {
				everAdjusted = true
			}
			if !isAdjusted {
				break
			}
			currPosition = adjustedPosition
		}

		currPlayer.movePiece(currPosition, everAdjusted)
		d.players = append(d.players[1:len(d.players)], currPlayer)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	cmds := []string{
		"9",
		"62 5",
		"33 6",
		"49 9",
		"88 16",
		"41 20",
		"56 53",
		"98 64",
		"93 73",
		"95 75",
		"8",
		"2 37",
		"27 46",
		"10 32",
		"51 68",
		"61 79",
		"65 84",
		"71 91",
		"81 100",
		"2",
		"Gaurav",
		"Sagar",
	}

	game := &game{}
	game.init(cmds)

	game.start()

}
