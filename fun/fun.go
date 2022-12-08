package main

import (
	"fmt"
	"time"
)

func main() {
	defer LogWithCtx("aaaaaaaaa")
	defer LogWithCtx("aaaaaaaaa1")()

}

// LogWithCtx 耗时统计函数 至BsLog.Debug
func LogWithCtx(desc string) func() {
	start := time.Now()
	return func() {
		fmt.Println(desc)
		tc := time.Since(start).Milliseconds()
		fmt.Println("************************", tc)
	}
}
