package main

import (
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"github.com/jodealter/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (b *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router handle...")
	fmt.Println("recv from client :msgID:", request.GetMsgId(), ",   msgData", request.GetData())
	err := request.GetConnection().SendMsg(1, []byte("ping ping ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	s := znet.NewServr("zinxv0.3")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
