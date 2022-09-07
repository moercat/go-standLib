package zrouter

import (
	"container/ring"
	"fmt"
	"github.com/aceld/zinx/ziface"
	"github.com/aceld/zinx/zlog"
	"github.com/aceld/zinx/znet"
	"github.com/golang/protobuf/proto"
	"go-standLib/game/pkg/proto/api"
	"math/rand"
	"strconv"
	"sync"
	"time"
)

var (
	Win   = ring.New(2)
	Close bool
	so    sync.Once
)

func init() {
	go so.Do(Cobra)
}

func Cobra() {
	t := time.NewTicker(2 * time.Second)
	r := rand.New(rand.NewSource(time.Now().Unix()))
	Win.Next().Value = r.Intn(100)

	for range t.C {
		if Win.Len() <= 10 {
			a := ring.New(1)
			a.Value = r.Intn(100)
			Win.Link(a)
		}
		Close = false
	}
}

type GuessRouter struct {
	znet.BaseRouter
}

func (this *GuessRouter) Handle(request ziface.IRequest) {
	zlog.Debug("now win number: ", Win, "Close", Close)
	//先读取客户端的数据，再回写ping...ping...ping
	zlog.Debug("recv from client : msgId=", request.GetMsgID(), ", data=", string(request.GetData()))

	msg := &api.Talk{}
	err := proto.Unmarshal(request.GetData(), msg)
	if err != nil {
		fmt.Println("Talk Unmarshal error ", err)
		return
	}

	var (
		winContent string
	)
	num, _ := strconv.Atoi(msg.Content)

	win := Win.Next().Value.(int)

	if !Close {
		if num == win {
			winContent = fmt.Sprintf("赢了,数字为%d", num)
			Win.Unlink(1)
			Close = true
		} else if num < win {
			winContent = fmt.Sprintf("太小了,数字比%d更大喔", num)
		} else if num > win {
			winContent = fmt.Sprintf("太大了,数字比%d更小喔", num)
		}
	} else {
		winContent = "结束了,等下一轮噢"
	}

	err = request.GetConnection().SendBuffMsg(2, []byte(winContent))
	if err != nil {
		zlog.Error(err)
	}
}
