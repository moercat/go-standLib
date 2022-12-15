package main

import (
	ping1 "go-standLib/plugin/pingo"
	"log"
	"time"
)

func main() {

	tricker := time.NewTicker(1 * time.Second)

	for range tricker.C {
		resp()
	}

}

func resp() {
	// 从创建的可执行文件中创建一个新插件。通过 TCP 连接到它
	p := ping1.NewPlugin("tcp", "plugin/hello-world/hello-world")
	// 启动插件
	p.Start()
	// 使用完插件后停止它
	defer p.Stop()

	var resp string

	// 从先前创建的对象调用函数
	if err := p.Call("MyPlugin.SayHello", "Go developer", &resp); err != nil {
		log.Print(err)
	} else {
		log.Print(resp)
	}
}
