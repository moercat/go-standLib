package do_test

import (
	"io"
	"io/ioutil"
	"log"
	"strings"
	"testing"
)

var str string

func init() {
	for i := 0; i < 100; i++ {
		str += "Go is a general-purpose language designed with systems programming in mind."
	}
}

func BenchmarkName_Iot(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(str)
		_, err := ioutil.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
	}
}

func BenchmarkName_Io(b *testing.B) {
	for i := 0; i < b.N; i++ {
		r := strings.NewReader(str)
		_, err := io.ReadAll(r)
		if err != nil {
			log.Fatal(err)
		}
	}
}
