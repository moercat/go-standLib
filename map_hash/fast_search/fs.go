package fast_search

import (
	"strconv"
)

var (
	tmpFenceHash  = NewStr()
	tmpFence2Hash = NewStr()
	tmpFence3Hash = NewStr()
	tmpFence4Hash = NewStr()
)

func Hash() {
	tmpFenceHash.SetNumberOfReplicas(20)
	tmpFence2Hash.SetNumberOfReplicas(20)
	tmpFence3Hash.SetNumberOfReplicas(20)
	tmpFence4Hash.SetNumberOfReplicas(20)
	for i := 0; i < 20; i++ {
		tmpFenceHash.Add(strconv.Itoa(i)+"lt-henan-xuchang-sn7-172-31-134-35", 50)
		tmpFence2Hash.Add(strconv.Itoa(i)+"lt-henan-kaifeng-sn8-172-31-121-45", 50)
		tmpFence3Hash.Add(strconv.Itoa(i)+"lt-henan-kaifeng-sn8-172-31-121-55", 50)
		tmpFence4Hash.Add(strconv.Itoa(i)+"lt-henan-xinyang-sn17-172-31-76-203", 50)
	}
	tmpFenceHash.Finally()
	tmpFence2Hash.Finally()
	tmpFence3Hash.Finally()
	tmpFence4Hash.Finally()
}

var get1 string
var idx int

func Get() {
	get, i, _ := tmpFenceHash.Get("aaa")
	//fmt.Println(get, i)

	get, i, _ = tmpFence2Hash.Get("aaa")
	//fmt.Println(get, i)

	get, i, _ = tmpFence3Hash.Get("aaa")
	//fmt.Println(get, i)

	get, i, _ = tmpFence4Hash.Get("aaa")
	//fmt.Println(get, i)

	get1 = get
	idx = i
}

func GetWithIdx() {
	get, i, _ := tmpFenceHash.Get("aaa")
	//fmt.Println(get, i)

	get, i, _ = tmpFence2Hash.GetWithIdx("aaa", i)
	//fmt.Println(get, i)

	get, i, _ = tmpFence3Hash.GetWithIdx("aaa", i)
	//fmt.Println(get, i)

	get, i, _ = tmpFence4Hash.GetWithIdx("aaa", i)
	//fmt.Println(get, i)

	get1 = get
	idx = i
}
