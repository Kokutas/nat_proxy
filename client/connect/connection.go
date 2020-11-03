package connect

import (
	"log"
	"nat_proxy/lib/network"
	"net"
)

// 客户端和远端的交互
func connRemote(remoteAddr string) (*net.TCPConn) {
	conn,err:= network.CreateTCPConnect(remoteAddr)
	if err != nil {
		log.Printf("与远端内网穿透服务代理 %v 建立连接失败 : %v\n",remoteAddr,err.Error())
	}
	return conn
}
// 客户端和本端交互：客户端和要暴露穿透的服务的连接
func connLocal(localAddr string)(*net.TCPConn)  {
	conn,err:= network.CreateTCPConnect(localAddr)
	if err != nil {
		log.Printf("与内网 %v 服务建立连接失败 : %v\n",localAddr,err.Error())
	}
	return conn
}

// 客户端和远端以及本端交互
func ConnLocalAndRemote(localAddr string,remoteAddr string)  {
	local:=connLocal(localAddr)
	remote:=connRemote(remoteAddr)
	// 如果远端和本端都连接成功
	if local != nil  && remote!=nil{
		network.JoinToConnect(local,remote)
	}else{
		if local != nil {
			local.Close()
		}
		if remote != nil {
			remote.Close()
		}
	}
}