package consistent

import "strconv"

var (
	_consistentMap    = New()
	_consistentStrMap = NewStr()
)

func Init() {
	consistentMap := New()
	consistentMap.SetNumberOfReplicas(20)
	var node []Node
	for i := 0; i < 20; i++ {
		var elt string
		for j := 0; j < 50; j++ {
			elt += strconv.Itoa(i)
		}
		node = append(node, Node{
			Name:   elt,
			Weight: 1000,
		})
	}
	consistentMap.Add(node)
	_consistentMap = consistentMap
}

func Get(str string) (string, error) {

	return _consistentMap.Get(str)
}

func InitStr() {
	consistentStrMap := NewStr()
	consistentStrMap.SetNumberOfReplicas(20)
	for i := 0; i < 20; i++ {
		var elt string
		for j := 0; j < 50; j++ {
			elt += strconv.Itoa(i)
		}
		consistentStrMap.Add(elt, 1000)
	}
	consistentStrMap.Finally()
	_consistentStrMap = consistentStrMap
}

func GetStr(str string) (string, error) {

	return _consistentStrMap.Get(str)
}
