package search

import (
	"math"

	"github.com/brice-allen/csci4202/utils"
)

/////////END STACK DEFs
type State struct {
	Board    [][]int
	NumMoves int
	Parent   *State
	LastMove Direction
	Distance int
}

type Direction int

const (
	None  Direction = 0
	Up    Direction = 1
	Down  Direction = 2
	Left  Direction = 3
	Right Direction = 4
)

const SizeX = 3
const SizeY = 3

//////COMPARES

func FindEmptyTile(board [][]int) (int, int) {
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			if board[y0][x0] == 0 {
				return x0, y0
			}
		}
	}
	return -1, -1
}

////////ACTIONS
func MoveUp(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyY > 0 {
		newBoard[emptyY-1][emptyX], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY-1][emptyX]
		return newBoard, emptyX, emptyY - 1
	}
	return board, emptyX, emptyY
}

func MoveDown(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyY < 2 {
		newBoard[emptyY+1][emptyX], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY+1][emptyX]
		return newBoard, emptyX, emptyY + 1
	}
	return board, emptyX, emptyY

}

func MoveLeft(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyX > 0 {
		newBoard[emptyY][emptyX-1], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY][emptyX-1]
		return newBoard, emptyX - 1, emptyY
	}
	return board, emptyX, emptyY

}

func MoveRight(board [][]int, emptyX int, emptyY int) ([][]int, int, int) {
	newBoard := make([][]int, 3)
	for i := 0; i < 3; i++ {
		newBoard[i] = make([]int, 3)
	}
	utils.CopySlice(newBoard, board)
	if emptyX < 2 {
		newBoard[emptyY][emptyX+1], newBoard[emptyY][emptyX] = newBoard[emptyY][emptyX], newBoard[emptyY][emptyX+1]
		return newBoard, emptyX + 1, emptyY
	}
	return board, emptyX, emptyY

}

//check possible moves to create children
func (s State) PossibleMoves() []State {
	x, y := FindEmptyTile(s.Board)
	goal := utils.FillGoal()
	moves := make([]State, 0, 4)

	if y > 0 {
		newBoard, _, _ := MoveUp(s.Board, x, y)
		moves = append(moves, State{Board: newBoard, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Up, Distance: ManhattanDistance(newBoard, goal)})
	}
	if y < 2 {
		newBoard, _, _ := MoveDown(s.Board, x, y)
		moves = append(moves, State{Board: newBoard, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Down, Distance: ManhattanDistance(newBoard, goal)})
	}
	if x > 0 {
		newBoard, _, _ := MoveLeft(s.Board, x, y)
		moves = append(moves, State{Board: newBoard, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Left, Distance: ManhattanDistance(newBoard, goal)})
	}
	if x < 2 {
		newBoard, _, _ := MoveRight(s.Board, x, y)
		moves = append(moves, State{Board: newBoard, NumMoves: s.NumMoves + 1, Parent: &s, LastMove: Right, Distance: ManhattanDistance(newBoard, goal)})
	}
	return moves

}

////BEGINING & END
func NewState(board [][]int, goal [][]int) State {
	return State{Board: board, Distance: ManhattanDistance(board, goal)}
}

func (s State) IsGoal(goal [][]int) bool {
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			if s.Board[y0][x0] != goal[y0][x0] {
				return false
			}
		}
	}
	return true
}

/////HEURISTIC METHODS
func ManhattanDistance(board [][]int, goal [][]int) int {
	var dx, dy int
	var sum int
	sum = 0
	for y0 := 0; y0 < 3; y0++ {
		for x0 := 0; x0 < 3; x0++ {
			for y1 := 0; y1 < 3; y1++ {
				for x1 := 0; x1 < 3; x1++ {
					if board[y0][x0] == goal[y1][x1] && board[y0][x0] != 0 {
						dx = int(math.Abs(float64(x0 - x1)))
						dy = int(math.Abs(float64(y0 - y1)))
						sum = sum + dx + dy
						//fmt.Printf("x0: %d, y0:%d, sta: %d,,y0:%d,x0:%d,goa:%d\n", x0, y0, board[y0][x0], x1, y1, goal[y1][x1])
						//fmt.Printf("dx:%d,dy:%d,sum:%d\n", dx, dy, sum)
					}
				}
			}

		}
	}
	return sum
}
