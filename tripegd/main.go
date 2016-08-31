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
	board := tripeg.BuildBoard(0)
	fmt.Println(board)
	fmt.Println("----------")
	board.Solve()
	for _, m := range board.MoveLog {
		fmt.Println(m)
	}
}
