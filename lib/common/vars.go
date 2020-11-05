package common

import (
	"net"
	"strconv"
	"time"
)

// 变量定义
var (
	// 内网穿透服务器公网IP
	serverIP = net.IPv4(0, 0, 0, 0)
	// 内网穿透服务器控制端口
	serverCtrlPort = 8888
	// 内网穿透服务器隧道端口
	serverTunnelProt = 9999
	// 内网穿透服务器web访问端口
	serverWebPort = 8080
	// 内网穿透服务器控制地址
	ServerAddress = serverIP.String()+":"+strconv.Itoa(serverCtrlPort)

	// 客户端连接服务器失败重试时间间隔
	DialFailedRetryTime = time.Second*3
)
