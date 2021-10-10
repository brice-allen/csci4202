package main

import (
	"fmt"
	"github.com/brice-allen/csci4202/astar"
	"github.com/brice-allen/csci4202/bfs"
	"github.com/brice-allen/csci4202/search"
	"github.com/brice-allen/csci4202/utils"
	"os"
	"strings"
	"time"
)

func main() {
	utils.PrintMemUsage()
	args := os.Args
	if len(args) != 2 {
		fmt.Println("Usage: hw01 bfs or hw01 astar!")
		return
	}
	var selection int
	if strings.Compare(args[1], "bfs") == 0 {
		fmt.Println("bfs")
		selection = 1
	} else if strings.Compare(args[1], "astar") == 0 {
		fmt.Println("astar")
		selection = 2
	} else {
		fmt.Println("Usage: hw01 bfs or hw01 astar!")
		return
	}
	fmt.Println("Starting...", time.Now())
	startTime := time.Now()
	goal := make([][]int, 9)
	goal = utils.FillGoal()
	board := utils.InputParser("input.txt")
	if board == nil {
		return
	}
	utils.StatePrinter(board)
	startState := search.NewState(board, goal)
	var solution *search.State
	var frontier, expanded int
	if selection == 1 {
		solution, frontier, expanded = bfs.Solve(startState, goal)
	} else if selection == 2 {
		solution, frontier, expanded = astar.Solve(startState, goal)
	} else {
		fmt.Println("Usage: hw01 bfs or hw01 astar!")
		return
	}
	if solution == nil {
		fmt.Println("Çözülemedi...")
		fmt.Println("Frontier'e Giren Düğüm Sayısı:", frontier)
		fmt.Println("Frontier'den Çıkan Düğüm Sayısı:", expanded)
		fmt.Println("End...", time.Now())
		return
	}
	s := solution
	steps := make([]*search.State, solution.NumMoves)
	for i := 0; i < solution.NumMoves; i++ {
		steps[i] = s
		s = s.Parent
	}
	for i := len(steps)/2 - 1; i >= 0; i-- {
		opp := len(steps) - 1 - i
		steps[i], steps[opp] = steps[opp], steps[i]
	}
	for _, next := range steps {
		utils.StatePrinter(next.Board)
	}

	fmt.Println("Çözüm Maliyeti:", solution.NumMoves)
	fmt.Println("Frontier'e Giren Düğüm Sayısı:", frontier)
	fmt.Println("Frontier'den Çıkan Düğüm Sayısı:", expanded)
	fmt.Println("End...", time.Now())
	elapsed := time.Since(startTime)
	fmt.Println("Process took:", elapsed)
	utils.PrintMemUsage()
	return
}
