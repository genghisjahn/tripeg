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
	tripeg.BuildBoard()
}
