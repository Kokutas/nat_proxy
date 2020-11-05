package network

import (
	"fmt"
	"log"
	"net"
)

// 网卡/网络适配器相关信息

// 网卡/网络适配器信息结构体
type Adaptor struct {
	// 网卡/网络适配器序号
	Index int
	// 网卡/适配器名称
	Name string
	// 网卡/网络适配器最大传输单元MTU
	Mtu int
	// 网卡/网络适配器MAC地址
	Mac string
	// 网卡/网络适配器Flags标记
	Flags string
	// 网卡/网络适配器IPv4
	IPv4 net.IP
	// 网卡/网络适配器IPv4掩码位
	IPv4MaskCount int
	// 网卡/网络适配器IPv4掩码总位数
	IPv4MaskTotal int
	// 网卡/网络适配器IPv6
	IPv6 net.IP
	// 网卡/网络适配器IPv6掩码位
	IPv6MaskCount int
	// 网卡/网络适配器IPv6掩码总位数
	IPv6MaskTotal int
}

// 根据IP获取网卡/网络适配器信息
func Adaptors(ip net.IP) (*Adaptor, error) {
	addresss, err := net.ResolveIPAddr("", ip.String())
	if err != nil {
		return nil, err
	}
	IP := net.ParseIP(addresss.String())
	var adaptor *Adaptor
	adaptors, err := adaptors()
	if err != nil {
		return nil, err
	}
	for _, adp := range adaptors {
		if adp.IPv4.Equal(IP) {
			adaptor = adp
			return adaptor, nil
		} else if adp.IPv6.Equal(IP) {
			adaptor = adp
			return adaptor, nil
		}
	}
	return nil, fmt.Errorf("Not fond adaptor of ip = \"%v\".\n", ip)
}

// 获取所有的网卡/网络适配器信息
func adaptors() ([]*Adaptor, error) {
	// 获取所有的网卡/网络适配器信息
	ifaces, err := net.Interfaces()
	if err != nil {
		return nil, err
	}
	adaptors := make([]*Adaptor, 0)
LAB:
	// 遍历所有的网卡
	for _, iface := range ifaces {
		// 如果状态标记不是UP的直接执行下一个
		if iface.Flags&net.FlagUp == 0 {
			log.Printf("The adaptor \"%v\" Flags is not up.\n", iface.Name)
			continue LAB
		}
		addresses, err := iface.Addrs()
		if err != nil {
			log.Printf("The adaptor \"%v\" get address failed.\n", iface.Name)
			continue LAB
		}
		// 如果地址切片的长度小于2执行下一个
		if len(addresses) < 2 {
			log.Printf("The adaptor \"%v\" address slice < 2.\n", iface.Name)
			continue LAB
		}

		ipv4 := net.IP{}
		ipv6 := net.IP{}
		ipv4MaskCount := 0
		ipv6MaskCount := 0
		ipv4MaskTotal := 0
		ipv6MaskTotal := 0
		for _, address := range addresses {
			ip, ipnet, err := net.ParseCIDR(address.String())
			if err != nil {
				log.Printf("The adaptor \"%v\" address \"%v\" parse cidr failed.\n", iface.Name, address.String())
				continue
			}
			// 如果是本地回环
			if ipnet.IP.IsLoopback() {
				log.Printf("The adaptor \"%v\" address \"%v\" is loopbak address.\n", iface.Name, ip)
				continue LAB
			}
			if ip.To4() != nil {
				ipv4MaskCount, ipv4MaskTotal = ipnet.Mask.Size()
				ipv4 = ip.To4()
			} else if ip.To16() != nil {
				ipv6MaskCount, ipv6MaskTotal = ipnet.Mask.Size()
				ipv6 = ip.To16()
			}
		}
		if ipv4 == nil || ipv6 == nil {
			continue LAB
		}
		adaptor := &Adaptor{
			Index:         iface.Index,
			Name:          iface.Name,
			Mtu:           iface.MTU,
			Mac:           iface.HardwareAddr.String(),
			Flags:         iface.Flags.String(),
			IPv4:          ipv4,
			IPv4MaskCount: ipv4MaskCount,
			IPv4MaskTotal: ipv4MaskTotal,
			IPv6:          ipv6,
			IPv6MaskCount: ipv6MaskCount,
			IPv6MaskTotal: ipv6MaskTotal,
		}
		adaptors = append(adaptors, adaptor)
	}
	return adaptors, nil
}
