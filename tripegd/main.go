package main

import (
	"crypto/md5"
	"fmt"
	"os"
	"strconv"

	"github.com/garyburd/redigo/redis"
	"github.com/genghisjahn/tripeg"
	"github.com/genghisjahn/xlog"
)

func main() {
	conn, err := redis.Dial("tcp", "localhost:6379")
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	solutions := map[string]string{}
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
	sCount := 0
	for s := 0; s < 5000; s++ {
		board, err := tripeg.BuildBoard(rows, empty)
		if err != nil {
			xlog.Error.Println(err)
			return
		}
		board.Solve()
		// for k, c := range board.MoveChart {
		// 	fmt.Println("Move:", k+1, fmt.Sprintf("%s", c))
		// }
		sol := ""
		for _, c := range board.MoveChart {
			sol += fmt.Sprintf("%s", c)
		}
		data := []byte(sol)
		sHash := fmt.Sprintf("%x", md5.Sum(data))
		if _, ok := solutions[sHash]; !ok {
			sCount++
			// fmt.Println(sCount, sHash)
			solutions[sHash] = fmt.Sprintf("%s", sol)
			conn.Do("set", sHash, sol)
		}
		if s%100 == 0 {
			fmt.Println(s, sCount)
		}
	}
	fmt.Println(len(solutions))
}
