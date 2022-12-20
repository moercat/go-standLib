package wr

import (
	"fmt"
	wr "github.com/mroth/weightedrand"
	wr2 "github.com/mroth/weightedrand/v2"
	"testing"
)

func BenchmarkName_wr(b *testing.B) {

	var r []wr.Choice

	for i := 0; i < 10; i++ {
		r = append(r, wr.NewChoice(i, uint(i)))
	}

	chooser, err := wr.NewChooser(r...)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		_ = chooser.Pick().(int)

	}
}

func BenchmarkName_wr2(b *testing.B) {
	var r []wr2.Choice[int, uint]

	for i := 0; i < 10; i++ {
		r = append(r, wr2.NewChoice[int, uint](i, uint(i)))
	}

	chooser, err := wr2.NewChooser(r...)
	if err != nil {
		fmt.Println(err)
	}

	for i := 0; i < b.N; i++ {
		_ = chooser.Pick()
	}
}
