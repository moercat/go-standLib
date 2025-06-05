package spilt

import "strings"

var a = "a'a'a'a'b'a'a'a'a'a"

func S() []string {
	return strings.Split(a, "'")
}

func SN() []string {
	return strings.SplitN(a, "'", 3)
}

func C() bool {
	return strings.Contains(a, "b")
}

func CR() bool {
	return strings.ContainsRune(a, 'b')
}
