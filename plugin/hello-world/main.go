// 创建新的二进制文件
package main

import (
	ping1 "go-standLib/plugin/pingo"
)

// 创建要导出的对象
type MyPlugin struct{}

// 导出的方法，带有rpc签名
func (p *MyPlugin) SayHello(name string, msg *string) error {
	*msg = "Hello, " + name + "1"
	return nil
}

func main() {
	plugin := &MyPlugin{}

	// 注册要导出的对象
	ping1.Register(plugin)
	// 运行主程序
	ping1.Run()
}
