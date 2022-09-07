package main

import (
	"fmt"
	"github.com/golang/protobuf/proto"
	"go-standLib/game/pkg/proto/api"
	"io"
	"net"
	"time"

	"github.com/aceld/zinx/znet"
)

/*
	模拟客户端
*/
func main() {

	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start err, exit!", err)
		return
	}

	//发封包message消息
	dp := znet.NewDataPack()

	go func() {
		for {
			var (
				str []byte
				p   = &api.Talk{}
			)
			_, err = fmt.Scan(&p.Content)
			if err != nil {
				return
			}

			if str, err = proto.Marshal(p); err != nil {
				return
			}

			msgs, _ := dp.Pack(znet.NewMsgPackage(2, str))
			_, err = conn.Write(msgs)
			if err != nil {
				fmt.Println("write error err ", err)
				return
			}
		}
	}()

	for {
		//先读出流中的head部分
		headData := make([]byte, dp.GetHeadLen())
		_, err = io.ReadFull(conn, headData) //ReadFull 会把msg填充满为止
		if err != nil {
			fmt.Println("read head error")
			break
		}
		//将headData字节流 拆包到msg中
		msgHead, err := dp.Unpack(headData)
		if err != nil {
			fmt.Println("server unpack err:", err)
			return
		}

		if msgHead.GetDataLen() > 0 {
			//msg 是有data数据的，需要再次读取data数据
			msg := msgHead.(*znet.Message)
			msg.Data = make([]byte, msg.GetDataLen())

			//根据dataLen从io中读取字节流
			_, err := io.ReadFull(conn, msg.Data)
			if err != nil {
				fmt.Println("server unpack data err:", err)
				return
			}

			fmt.Println("==> Test Router:[Ping] Recv Msg: ID=", msg.ID, ", len=", msg.DataLen, ", data=", string(msg.Data))
		}

		time.Sleep(2 * time.Second)
	}
}
