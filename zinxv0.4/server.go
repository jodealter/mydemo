package main

import (
	"fmt"
	"github.com/jodealter/zinx/ziface"
	"github.com/jodealter/zinx/znet"
)

type PingRouter struct {
	znet.BaseRouter
}

func (b *PingRouter) PreHandle(request ziface.IRequest) {
	fmt.Println("call router Prehandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("before ping..."))
	if err != nil {
		fmt.Println("call ping prehandle error : ", err)
	}

}

func (b *PingRouter) Handle(request ziface.IRequest) {
	fmt.Println("call router handle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("ping ping ping..."))
	if err != nil {
		fmt.Println("call ping posthandle error : ", err)
	}

}

func (b *PingRouter) PostHandle(request ziface.IRequest) {
	fmt.Println("call router Posthandle...")
	_, err := request.GetConnection().GetTCPConnection().Write([]byte("post ping..."))
	if err != nil {
		fmt.Println("call ping posthandle error : ", err)
	}

}

func main() {
	s := znet.NewServr("zinxv0.3")
	s.AddRouter(&PingRouter{})
	s.Serve()
}
