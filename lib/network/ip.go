package network

import (
	"net"
)

// ip 相关的

// 1.获取本机公网出口网卡的ip地址--本机内网IP
func LocalIP() (net.IP, error) {
	conn, err := net.Dial("tcp", "baidu.com:80")
	if err != nil {
		return nil, err
	}
	_ = conn.Close()
	addr, err := net.ResolveTCPAddr("tcp", conn.LocalAddr().String())
	if err != nil {
		return nil, err
	}
	return addr.IP, nil
}
