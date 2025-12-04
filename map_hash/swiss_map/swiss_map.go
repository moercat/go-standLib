package swiss_map

var m = make(map[int]string)

func Init() {
	for i := 0; i < 10000; i++ {
		m[i] = "1"
	}
}

func GetMap() string {
	for i := 0; i < 100000; i++ {
		m[i] = "1"
	}
	for i := 0; i < 3000; i++ {
		delete(m, i)
	}
	for i := 0; i < 10000; i++ {
		m[i] = "1"
	}
	return "1"
}
