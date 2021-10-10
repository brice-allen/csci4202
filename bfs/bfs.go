package bfs

import (
	_ "fmt"
	"github.com/brice-allen/csci4202/search"
	"github.com/brice-allen/csci4202/utils"
	//"time"
)

func Solve(start search.State, goal [][]int) (*search.State, int, int) {
	var frontier, expanded int
	states := make(map[string]search.State)
	var stack []string
	key := utils.BoardStringer(start.Board)
	states[key] = start
	stack = append(stack, key)
	for len(stack) != 0 {
		//time.Sleep(100 * time.Millisecond)
		currentItem := stack[0]
		stack[0] = ""
		stack = stack[1:]
		current := states[currentItem]
		expanded++
		//utils.StatePrinter(current.Board)
		//fmt.Println(current.Distance)
		if current.IsGoal(goal) {
			//solved
			return &current, frontier, expanded
		}
		for _, next := range current.PossibleMoves() {
			//implement key of state to keep in heap
			key := utils.BoardStringer(next.Board)
			if _, exists := states[key]; !exists {
				states[key] = next
				stack = append(stack, key)
				frontier++
			}
		}
	}
	return nil, frontier, expanded
}
