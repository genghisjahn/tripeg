package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/genghisjahn/tripeg"
	"github.com/genghisjahn/xlog"
)

func main() {
	if err := xlog.New(xlog.Infolvl); err != nil {
		panic(err)
	}
	argsWithoutProg := os.Args[1:]
	xlog.Info.Println("Tripeg Main")
	empty := 0
	rows := 5
	if len(argsWithoutProg) > 0 {
		v, vErr := strconv.Atoi(argsWithoutProg[0])
		if vErr != nil {
			xlog.Error.Println(vErr)
			return
		}
		empty = v
	}
	if len(argsWithoutProg) > 1 {
		v, vErr := strconv.Atoi(argsWithoutProg[1])
		if vErr != nil {
			xlog.Error.Println(vErr)
			return
		}
		rows = v
	}
	board, err := tripeg.BuildBoard(rows, empty)
	if err != nil {
		xlog.Error.Println(err)
		return
	}
	board.Solve()

	for k, c := range board.MoveChart {
		fmt.Println("Move:", k+1, fmt.Sprintf("%s", c))
	}
}
