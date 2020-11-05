package main

import (
	"log"
	"nat_proxy/lib/common"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
	psignal := make(chan os.Signal, 1)
	// ctrl+c->SIGINT, kill -9 -> SIGKILL
	signal.Notify(psignal, syscall.SIGINT, syscall.SIGKILL)
	go DialServer(common.ServerAddress)
	<-psignal
	log.Println("Bye~")
}
func DialServer(address string) {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		log.Fatalf("服务端地址解析 %v 错误\n", address)
	}
	// 连接服务端
	conn, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		log.Println("服务端断线，重新连接...", err)
		time.Sleep(common.DialFailedRetryTime)
		go DialServer(address)
		return
	}
	defer conn.Close()
	log.Printf("成功连接服务器 %v \n",address)

}
