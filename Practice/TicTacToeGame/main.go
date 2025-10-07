// Question - https://workat.tech/machine-coding/practice/design-tic-tac-toe-smyfi9x064ry

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Game interface {
	initializeGame()
	startGame()
	endGame()
}

type TicTacToe struct {
	Board   Board
	Players []Player
}

func (d *TicTacToe) initializeGame() {
	// initialize the board
	board := &ThreeByThreeBoard{}
	board.addNewBoard()
	d.Board = board

	board.printBoard()
}

func (d *TicTacToe) startGame() {
	var currPlayer Player
	row, column := -1, -1
	for !d.Board.isWinnerNew(currPlayer, row, column) {
		// check if it's a tie
		if d.Board.isTie() {
			fmt.Println("It's a tie. No one won")
			return
		}

		// get the current player
		currPlayer = d.Players[0]

		// get the positions
		rowAndColumn := getInput("[" + currPlayer.Name + "][" + currPlayer.Piece.Type + "] Enter row and column, indexed 0 (Eg: 0 1)")
		rowAndColumnArr := strings.Fields(rowAndColumn)
		row, _ = strconv.Atoi(rowAndColumnArr[0])
		column, _ = strconv.Atoi(rowAndColumnArr[1])

		// make the move
		isMoveValid := d.Board.move(currPlayer, row, column)
		if !isMoveValid {
			continue
		}

		// print the board
		d.Board.printBoard()
		// send the player back to queue
		d.Players = append(d.Players[1:], currPlayer)
	}
	fmt.Println("Player: ", currPlayer.Name, "(", currPlayer.Piece.Type, ")won")
	return
	// return the result
}

func (d *TicTacToe) endGame() {
	fmt.Println("Exiting the game")
	os.Exit(1)
}

type Board interface {
	addNewBoard()
	move(Player, int, int) bool
	isTie() bool
	printBoard()
	isWinner(Player) bool
	isWinnerNew(Player, int, int) bool
}

type ThreeByThreeBoard struct {
	Pieces           [][]Piece
	RowCountMap      map[string][3]int
	ColumnCountMap   map[string][3]int
	DiagonalCountMap map[string][2]int
}

func (d *ThreeByThreeBoard) addNewBoard() {
	var piecesMatrix [][]Piece
	for i := 0; i < 3; i++ {
		var piecesRow []Piece
		for j := 0; j < 3; j++ {
			p := Piece{Type: "--"}
			piecesRow = append(piecesRow, p)
		}
		piecesMatrix = append(piecesMatrix, piecesRow)
	}
	d.Pieces = piecesMatrix

}

func (d *ThreeByThreeBoard) printBoard() {
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			fmt.Print(d.Pieces[i][j].Type + "  ")
		}
		fmt.Print("\n")
	}
	fmt.Println("----------------")
}

/*
This algo can be optimised
*/
func (d *ThreeByThreeBoard) isWinner(player Player) bool {
	if player.Name == "" {
		return false
	}
	for i := 0; i < 3; i++ {
		if d.Pieces[i][1].Type == "--" {
			continue
		}
		if d.Pieces[i][0].Type == d.Pieces[i][1].Type && d.Pieces[i][1].Type == d.Pieces[i][2].Type {
			fmt.Println(player.Name, "won horizontally")
			return true
		}
	}

	for i := 0; i < 3; i++ {
		if d.Pieces[1][i].Type == "--" {
			continue
		}
		if d.Pieces[0][i].Type == d.Pieces[1][i].Type && d.Pieces[1][i].Type == d.Pieces[2][i].Type {
			fmt.Println(player.Name, "won vertically")
			return true
		}
	}

	if d.Pieces[0][0].Type != "--" && d.Pieces[0][0].Type == d.Pieces[1][1].Type && d.Pieces[0][0].Type == d.Pieces[2][2].Type {
		fmt.Println(player.Name, "won diagonally")
		return true
	}

	if d.Pieces[0][2].Type != "--" && d.Pieces[0][2].Type == d.Pieces[1][1].Type && d.Pieces[0][2].Type == d.Pieces[2][0].Type {
		fmt.Println(player.Name, "won diagonally")
		return true
	}
	return false
}

func (d *ThreeByThreeBoard) isWinnerNew(player Player, row, column int) bool {
	if row == -1 || column == -1 {
		return false
	}

	if d.RowCountMap[player.Piece.Type][row] == 3 {
		fmt.Println(player.Name, "won horizontally")
		return true
	}

	if d.ColumnCountMap[player.Piece.Type][column] == 3 {
		fmt.Println(player.Name, "won vertically")
		return true
	}

	if d.DiagonalCountMap[player.Piece.Type][0] == 3 {
		fmt.Println(player.Name, "won diagonally")
		return true
	}

	if d.DiagonalCountMap[player.Piece.Type][1] == 3 {
		fmt.Println(player.Name, "won diagonally")
		return true
	}
	return false
}

func (d *ThreeByThreeBoard) isTie() bool {
	var emptyPositionLeft int
	for i := 0; i < 3; i++ {
		for j := 0; j < 3; j++ {
			if d.Pieces[i][j].Type == "--" {
				emptyPositionLeft++
			}
		}
	}
	if emptyPositionLeft == 0 {
		return true
	}
	return false

}

func (d *ThreeByThreeBoard) move(player Player, row, column int) bool {
	// check if valid positions
	if row > 2 || column > 2 {
		fmt.Println("Invalid Move. Rows and column must be less than 2 in a 3*3 game")
		return false
	}
	// check if the position is already filled
	if d.Pieces[row][column].Type != "--" {
		fmt.Println("Invalid Move. This position is already taken by", d.Pieces[row][column].Type)
		return false
	}
	fmt.Println("Position[", row, "][", column, "] filled by", player.Piece.Type)
	d.Pieces[row][column].Type = player.Piece.Type

	if len(d.RowCountMap) == 0 {
		d.RowCountMap = make(map[string][3]int)
	}
	if len(d.ColumnCountMap) == 0 {
		d.ColumnCountMap = make(map[string][3]int)
	}
	if len(d.DiagonalCountMap) == 0 {
		d.DiagonalCountMap = make(map[string][2]int)
	}

	rowCounts := d.RowCountMap[player.Piece.Type]
	rowCounts[row]++
	d.RowCountMap[player.Piece.Type] = rowCounts

	columnCounts := d.ColumnCountMap[player.Piece.Type]
	columnCounts[column]++
	d.ColumnCountMap[player.Piece.Type] = columnCounts

	diagonalCounts := d.DiagonalCountMap[player.Piece.Type]
	if row == column {
		diagonalCounts[0]++
	}
	if 2-row == column {
		diagonalCounts[1]++
	}
	d.DiagonalCountMap[player.Piece.Type] = diagonalCounts

	return true
}

type Player struct {
	Name  string
	Piece Piece
}

type Piece struct {
	Type string
}

func main() {
	player1Details := "X Shashank"
	player2Details := "O Prakash"

	player1DetailsArr := strings.Fields(player1Details)
	player1Symbol := player1DetailsArr[0]
	player1Name := player1DetailsArr[1]
	player2DetailsArr := strings.Fields(player2Details)
	player2Symbol := player2DetailsArr[0]
	player2Name := player2DetailsArr[1]

	// initialize the game
	game := &TicTacToe{
		Players: []Player{
			Player{
				Name: player1Name,
				Piece: Piece{
					Type: player1Symbol,
				},
			},
			Player{
				Name: player2Name,
				Piece: Piece{
					Type: player2Symbol,
				},
			},
		},
	}
	game.initializeGame()
	game.startGame()
	game.endGame()

}

func getInput(label string) string {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print(label)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input) // remove newline
	return input
}
