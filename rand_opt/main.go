package main

import (
	opt "go-standLib/rand_opt/rand_opt"
	"net/http"
	_ "net/http/pprof"
)

func main() {

	go func() {
		_ = http.ListenAndServe("localhost:6060", nil)
	}()

	var r opt.RNG

	for i := 0; i < 10000000; i++ {
		//opt.MathRand()
		opt.Rng(r)
		//fmt.Println(i)
	}

	_ = http.ListenAndServe("127.0.0.1:18509", nil)

}
