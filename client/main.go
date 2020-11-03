package main

import (
	"bufio"
	"io"
	"log"
	"nat_proxy/client/connect"
	"nat_proxy/lib/network"
)

var (
	// 本机需要暴露的服务地址：需要进行内网穿透的地址
	localServerAddr = "127.0.0.1:8080"
	// 内网穿透中具有公网IP的ECS的ip
	remoteIP = "10.0.0.1"
	// 内网穿透中的ECS的内网穿透的服务控制通道：传递控制信息--注册到远端、心跳、出现新连接
	remoteControlAddr = remoteIP + ":9999"
	// 内网穿透中的ECS的内网穿透的服务隧道通道：建立转发隧道
	remoteServerAddr = remoteIP + ":8888"
)

func main() {
	// 和远端建立控制连接
	conn, err := network.CreateTCPConnect(remoteControlAddr)
	if err != nil {
		log.Fatalf("与远端控制 %v 建立连接失败: %v\n", remoteControlAddr, err)
	}
	log.Printf("与远端%v控制服务建立连接.\n",remoteControlAddr)
	// 创建读取器
	reader := bufio.NewReader(conn)
	for {
		content, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			log.Println(err)
		}
		// 当有新的连接信号出现时，创建一个新的连接--建立隧道
		if content == network.NewConnection+"\n" {
			go connect.ConnLocalAndRemote(localServerAddr,remoteServerAddr)
		}
	}
	log.Printf("与远端%v控制服务断开连接.\n",remoteControlAddr)
}
