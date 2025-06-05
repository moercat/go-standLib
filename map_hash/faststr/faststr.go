package faststr

import (
	"math/rand"
	"strconv"
)

var (
	m = make(map[int]string, 50)
	s = make([]string, 50)
)

func init() {
	for i := 0; i < 50; i++ {
		m[i] = strconv.Itoa(i)
	}

	for i := 0; i < 50; i++ {
		s[i] = strconv.Itoa(i)
	}

}

func PMap() string {
	return m[rand.Intn(50)]
}

func PSlice() string {
	return s[rand.Intn(50)]
}
