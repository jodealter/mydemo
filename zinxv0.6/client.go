package main

import (
	"fmt"
	"github.com/jodealter/zinx/znet"
	"io"
	"net"
	"time"
)

func main() {
	fmt.Println("client start...")
	time.Sleep(1 * time.Second)
	conn, err := net.Dial("tcp", "127.0.0.1:8999")
	if err != nil {
		fmt.Println("client start error exit : ", err)
		return
	}
	for {

		dp := znet.NewDataPack()
		binaryMsg, err := dp.Pack(znet.NewMessage(1, []byte("zinxv0.6 client test message")))
		if err != nil {
			fmt.Println("pack error : ", err)
			return
		}
		if _, err := conn.Write(binaryMsg); err != nil {
			fmt.Println("write to server error :", err)
			return
		}

		//服务器会给我们发送一msg个数据
		//先从流中读取头部数据，再接着读取剩余部分
		binaryHead := make([]byte, dp.GetHeadLen())

		if _, err := io.ReadFull(conn, binaryHead); err != nil {
			fmt.Println("read head error :", err)
			return
		}

		//将二进制拆包后的head放到msghead中
		msgHead, err := dp.UnPack(binaryHead)
		if err != nil {
			fmt.Println("unpack error : ", err)
			break
		}
		if msgHead.GetMsglen() > 0 {
			//进行二次读取
			msg := msgHead.(*znet.Message)
			data := make([]byte, msg.GetMsglen())
			if _, err := io.ReadFull(conn, data); err != nil {
				fmt.Println("read msg data error : ", err)
				return
			}
			msg.SetData(data)
			fmt.Println("recv server msg id:", msg.Id, "  len = ", msg.GetMsglen(), "  data = ", string(msg.GetData()))
		}
		time.Sleep(1 * time.Second)
	}
}
