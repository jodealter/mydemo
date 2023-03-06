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
	fmt.Println("call router ping handle...")
	fmt.Println("recv from client :msgID:", request.GetMsgId(), ",   msgData", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("ping ping ping..."))
	if err != nil {
		fmt.Println(err)
	}
}

type HelloRouter struct {
	znet.BaseRouter
}

func (h *HelloRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router hello handle...")
	fmt.Println("recv from client :msgID:", request.GetMsgId(), ",   msgData", string(request.GetData()))
	err := request.GetConnection().SendMsg(1, []byte("hello welcome come to jia..."))
	if err != nil {
		fmt.Println(err)
	}
}

// 创建链接之后的钩子函数
func DoConnectionBegin(conn ziface.IConnection) {
	fmt.Println("=====>DoConnectionBegin is Called...")
	if err := conn.SendMsg(202, []byte("DoConnection Begin...")); err != nil {
		fmt.Println(err)
	}
}

// 链接断开之后的钩子函数
func DoConnectionLost(conn ziface.IConnection) {
	fmt.Println("=====>DoConnectionLost is Called...")
	fmt.Println("[[[conn ID = ", conn.GetConnID(), " is Lost..]]]")
}

func main() {
	s := znet.NewServr("zinxv0.3")
	s.AddRouter(0, &PingRouter{})
	s.AddRouter(1, &HelloRouter{})
	s.SetOnConnStart(DoConnectionBegin)
	s.SetOnConnStop(DoConnectionLost)
	s.Serve()
}
