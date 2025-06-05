package live_templates

func ADD() {

	// 实时模版 两次循环 for2
	//for _, $VALUE$ := range $COLLECTION$ {
	//    for _, $VALUE2$ := range $VALUE$ {
	//        $END$
	//    }
	//}

	// example
	//var m = make(map[string][]string)
	//for _, strings := range m {
	//    for _, s := range strings {
	//
	//    }
	//}

	// 实时模版 用文件名做参数
	//logTag := "$VALUE$"
	// example
	//logTag := "2-add"

	// 实时模版 遍历复制的map
	// example
	//var m = make(map[string]int)
	//for s, i := range m {
	//
	//}
}

func ADD2() {

	// 环绕模版 更好的 make
	//var m = make(map[string]string, 0)
	//var strings = make([]string, 0, 0)
}
