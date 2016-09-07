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
	board := tripeg.BuildBoard(3)
	fmt.Println(board)
	board.Solve()
	fmt.Println("----------")
	for _, m := range board.MoveLog {
		fmt.Println(m)
	}
	fmt.Println("----------")

}
