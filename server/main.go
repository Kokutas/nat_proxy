package main

import (
	"log"
	"nat_proxy/lib/network"
	"net"
	"strconv"
	"sync"
	"time"
)

// 常量定义
const (
	// 与客户端建立连接的控制服务监听
	controlAddr = "0.0.0.0:9999"
	// 隧道地址：内网穿透服务与内网暴露的服务隧道，io.copy的监听,端口映射/绑定
	tunnelAddr = "0.0.0.0:8888"
	// 访问地址
	visitAddr = "0.0.0.8:7777"
)

// 全局变量定义
var (
	// 客户端连接记录
	clientConn *net.TCPConn
	// 连接池 : TODO:线程安全的map,map[string]interface{}--map[string]*ConnMatch
	connectionPool map[string]*ConnMatch
	// 连接池互斥锁
	connectionPoolLock sync.Mutex
)

type ConnMatch struct {
	// 接入连接池的时间
	addTime time.Time
	// 建立的连接
	accept *net.TCPConn
}

func main() {
	// 初始化连接池：初始设定为32
	connectionPool = make(map[string]*ConnMatch, 32)
	// 开启goroutine去创建控制通道监听，用以和客户端建联
	go createContrlChan()
}

// 创建控制通道，用以和客户端传递消息：心跳，创建新连接
func createContrlChan() {
	listener, err := network.CreateTCPListener(controlAddr)
	if err != nil {
		log.Fatal(err)
	}
	//defer listener.Close()
	log.Printf("创建监听控制通道服务%v\n", controlAddr)
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			log.Println(err)
			continue
		}
		log.Printf("新连接%v\n", conn.RemoteAddr().String())
		// TODO：只能允许一个客户端连接,如果多个，建议做个队列，比如slice的chan
		// 如果当前已经有一个客户端连接存在，则丢弃此连接
		if clientConn != nil {
			_ = conn.Close()
		} else {
			// 如果当前的客户端连接不存在，进行赋值记录，并开启心跳处理
			clientConn = conn
			go keepAlive()
		}
	}
}

// 和客户端保持连接：心跳机制
func keepAlive() {
	go func() {
		for {
			if clientConn == nil {
				return
			}
			// 发送心跳
			_, err := clientConn.Write([]byte(network.KeepAlive + "\n"))
			if err != nil {
				log.Printf("客户端%v已经断开连接\n", clientConn.RemoteAddr().String())
				clientConn.Close()
				return
			}
			// 3秒后继续
			time.Sleep(time.Second * 3)
		}
	}()
}

// 监听来自用户的请求:用户通过浏览器访问的请求--即发生一次访问被穿透的内网请求，7777--888-->内网的ip:8080
func acceptUserRequest() {
	listener, err := network.CreateTCPListener(visitAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		// 添加到连接池中
		addConnToPool(conn)
		// 发送消息给内网的客户端，通知建立隧道
		sendMessage(network.NewConnection)
	}
}

// 将访问的连接放到连接池中
func addConnToPool(conn *net.TCPConn) {
	connectionPoolLock.Lock()
	defer connectionPoolLock.Unlock()
	now := time.Now()
	connectionPool[strconv.FormatInt(now.UnixNano(), 10)] = &ConnMatch{
		addTime: now,
		accept:  conn,
	}
}

// 发送给客户端消息
func sendMessage(message string) {
	if clientConn == nil {
		log.Println("无连接中的客户端")
		return
	}
	_, err := clientConn.Write([]byte(message))
	if err != nil {
		log.Printf("向客户端发送消息%v异常：%v\n", message, err.Error())
	}
}

// 接受来自客户端的隧道建立请求并建立隧道
func acceptClientTunnelRequest() {
	listener, err := network.CreateTCPListener(tunnelAddr)
	if err != nil {
		log.Fatal(err)
	}
	defer listener.Close()
	for {
		conn, err := listener.AcceptTCP()
		if err != nil {
			continue
		}
		// 开启goroutine去执行建立隧道
		go establishTunnel(conn)
	}
}

// 建立隧道
func establishTunnel(tunnel *net.TCPConn) {
	// 锁定连接池
	connectionPoolLock.Lock()
	defer connectionPoolLock.Unlock()
	// 隧道关闭
	defer tunnel.Close()
	// 遍历连接池
	for key, connMatch := range connectionPool {
		if connMatch != nil {
			// 开启goroutine进行连接池中的连接和隧道请求进行io copy
			go network.JoinToConnect(connMatch.accept, tunnel)
			// 通过key--加入时间进行从连接池删除记录
			delete(connectionPool, key)
			return
		}
	}

}

// 清空连接池:10秒清空一次
func cleanConnectionPool() {
	for {
		connectionPoolLock.Lock()
		for key, connMatch := range connectionPool {
			// 超过10秒关闭连接，从连接池删除
			if time.Now().Sub(connMatch.addTime) > time.Second*10 {
				connMatch.accept.Close()
				delete(connectionPool, key)
			}
		}
		connectionPoolLock.Unlock()
		// 延迟5秒
		time.Sleep(time.Second*5)
	}
}
