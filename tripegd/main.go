package main

import (
	"github.com/genghisjahn/tripeg"
	"github.com/genghisjahn/xlog"
)

func main() {
	if err := xlog.New(xlog.Infolvl); err != nil {
		panic(err)
	}
	xlog.Info.Println("Tripeg Main")
	board := tripeg.BuildBoard()
	xlog.Info.Println(board)
}
