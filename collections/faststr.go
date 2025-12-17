package collections

import (
	"math/rand"
	"strconv"
)

func FastStrInit() (map[int]string, []string) {
	m := make(map[int]string, 50)
	s := make([]string, 50)

	for i := 0; i < 50; i++ {
		m[i] = strconv.Itoa(i)
	}

	for i := 0; i < 50; i++ {
		s[i] = strconv.Itoa(i)
	}

	return m, s
}

func PMap() string {
	m, _ := FastStrInit()
	return m[rand.Intn(50)]
}

func PSlice() string {
	_, s := FastStrInit()
	return s[rand.Intn(50)]
}
