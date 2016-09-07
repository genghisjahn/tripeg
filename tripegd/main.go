package main

import (
	"fmt"

	"github.com/genghisjahn/tripeg"
	"github.com/genghisjahn/xlog"
)

func main() {
	if err := xlog.New(xlog.Infolvl); err != nil {
		panic(err)
	}
	xlog.Info.Println("Tripeg Main")
	board, err := tripeg.BuildBoard(16)
	if err != nil {
		xlog.Error.Println(err)
		return
	}
	fmt.Println(board)
	board.Solve()
	for _, m := range board.MoveLog {
		fmt.Println(m)
	}
}
