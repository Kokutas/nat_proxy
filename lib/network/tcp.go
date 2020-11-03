package network

import (
	"io"
	"log"
	"net"
)

// 定义常量
const (
	// 保活标记
	KeepAlive = "KEEP_ALIVE"
	// 新连接状态字
	NewConnection = "NEW_CONNECTION"
)

// 新建TCP监听器
func CreateTCPListener(address string) (*net.TCPListener, error) {
	addr,err:= net.ResolveTCPAddr("tcp",address)
	if err != nil {
		return nil, err
	}
	// 监听TCP连接
	return net.ListenTCP("tcp", addr)
}

// 新建TCP连接
func CreateTCPConnect(address string) (*net.TCPConn, error) {
	addr, err := net.ResolveTCPAddr("tcp", address)
	if err != nil {
		return nil, err
	}
	return net.DialTCP("tcp",nil,addr)
}

// 加入到连接的方法
func JoinToConnect(local *net.TCPConn, remote *net.TCPConn) {
	// 开启goroutine: 从本地连接到远端
	go joinConnect(local, remote)
	// 开启goroutine: 从远端连接到本地
	go joinConnect(remote, local)
}

// 加入连接的方法: 其本质就是进行IO 的 Copy
func joinConnect(local *net.TCPConn, remote *net.TCPConn) {
	// 延迟关闭连接
	defer local.Close()
	defer remote.Close()
	_, err := io.Copy(local, remote)
	if err != nil {
		log.Printf("io.Copy Error:%v\n", err)
		return
	}
}
