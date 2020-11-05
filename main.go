package main

import (
	"bufio"
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io"
	"os/exec"
	"runtime"
)

/*
// TODO：链路追踪
// 路由？？
// https://jingyan.baidu.com/article/d169e1867b217f436711d856.html
// https://jingyan.baidu.com/article/870c6fc30abdf0b03fe4beb3.html
// 默认网关？？？
// https://ask.csdn.net/questions/1020627?utm_medium=distribute.pc_aggpage_search_result.none-task-ask_topic-2~all~first_rank_v2~rank_v28-6-1020627.nonecase&utm_term=golang%E8%8E%B7%E5%8F%96%E9%BB%98%E8%AE%A4%E7%BD%91%E5%85%B3&spm=1000.2123.3001.4430
func main() {
	//rifs, _ := nettest.RoutedInterface("ip", net.FlagUp | net.FlagBroadcast)
	//if rifs != nil {
	//	fmt.Println("Routed interface is ",rifs.HardwareAddr.String(),rifs.Name)
	//	fmt.Println(rifs.Addrs())
	//	fmt.Println("Flags are", rifs.Flags.String())
	//}
	// 执行cmd的命令
	cmd := exec.Command("route", "print", "-4")
	data, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s\n", data)
}
*/

type Charset string

const (
	UTF8    = Charset("UTF-8")
	GB18030 = Charset("GB18030")
)

func main() {
	// 操作系统判断
	switch runtime.GOOS {
	case "windows":
		//execCommand("route", []string{"print"})
		//execCommand("route", []string{"print","-4"})

		execCommand("route", []string{"print","-4","192*"})

	}

}
/*
ipconfig命令帮助
> ipconfig                       ... 显示信息
    > ipconfig /all                  ... 显示详细信息
    > ipconfig /renew                ... 更新所有适配器
    > ipconfig /renew EL*            ... 更新所有名称以 EL 开头
                                         的连接
    > ipconfig /release *Con*        ... 释放所有匹配的连接，
                                         例如“有线以太网连接 1”或
                                             “有线以太网连接 2”
    > ipconfig /allcompartments      ... 显示有关所有隔离舱的
                                         信息
    > ipconfig /allcompartments /all ... 显示有关所有隔离舱的
                                         详细信息
 */
// TODO：CMD 采集结果如下：
/*
[route print -4]
===========================================================================
接口列表
 15...00 ff 9d 78 0c b1 ......TAP-Windows Adapter V9
 23...f4 8e 38 f6 b8 63 ......Realtek PCIe GbE Family Controller
  4...02 00 4c 4f 4f 50 ......Npcap Loopback Adapter
  3...0a 00 27 00 00 03 ......VirtualBox Host-Only Ethernet Adapter
 18...bc a8 a6 b2 fc e2 ......Microsoft Wi-Fi Direct Virtual Adapter
  9...be a8 a6 b2 fc e1 ......Microsoft Wi-Fi Direct Virtual Adapter #2
 12...00 50 56 c0 00 01 ......VMware Virtual Ethernet Adapter for VMnet1
 20...00 50 56 c0 00 08 ......VMware Virtual Ethernet Adapter for VMnet8
 22...bc a8 a6 b2 fc e1 ......Intel(R) Dual Band Wireless-AC 3165
 13...bc a8 a6 b2 fc e5 ......Bluetooth Device (Personal Area Network)
  1...........................Software Loopback Interface 1
===========================================================================

IPv4 路由表
===========================================================================
活动路由:
网络目标        网络掩码          网关       接口   跃点数
          0.0.0.0          0.0.0.0      192.168.0.1     192.168.0.26     55
          0.0.0.0          0.0.0.0      192.168.8.2      192.168.8.1    291
          0.0.0.0        128.0.0.0       26.26.26.3       26.26.26.1      1
       26.26.26.0  255.255.255.248            在链路上        26.26.26.1    257
       26.26.26.1  255.255.255.255            在链路上        26.26.26.1    257
       26.26.26.7  255.255.255.255            在链路上        26.26.26.1    257
        127.0.0.0        255.0.0.0            在链路上         127.0.0.1    331
        127.0.0.1  255.255.255.255            在链路上         127.0.0.1    331
  127.255.255.255  255.255.255.255            在链路上         127.0.0.1    331
        128.0.0.0        128.0.0.0       26.26.26.3       26.26.26.1      1
      169.254.0.0      255.255.0.0            在链路上    169.254.75.113    291
      169.254.0.0      255.255.0.0            在链路上    169.254.254.97    281
   169.254.75.113  255.255.255.255            在链路上    169.254.75.113    291
   169.254.254.97  255.255.255.255            在链路上    169.254.254.97    281
  169.254.255.255  255.255.255.255            在链路上    169.254.75.113    291
  169.254.255.255  255.255.255.255            在链路上    169.254.254.97    281
      192.168.0.0    255.255.255.0            在链路上      192.168.0.26    311
     192.168.0.26  255.255.255.255            在链路上      192.168.0.26    311
    192.168.0.255  255.255.255.255            在链路上      192.168.0.26    311
      192.168.8.0    255.255.255.0            在链路上       192.168.8.1    291
      192.168.8.1  255.255.255.255            在链路上       192.168.8.1    291
    192.168.8.255  255.255.255.255            在链路上       192.168.8.1    291
     192.168.18.0    255.255.255.0            在链路上      192.168.18.1    281
     192.168.18.1  255.255.255.255            在链路上      192.168.18.1    281
   192.168.18.255  255.255.255.255            在链路上      192.168.18.1    281
        224.0.0.0        240.0.0.0            在链路上         127.0.0.1    331
        224.0.0.0        240.0.0.0            在链路上      192.168.18.1    281
        224.0.0.0        240.0.0.0            在链路上      192.168.0.26    311
        224.0.0.0        240.0.0.0            在链路上    169.254.75.113    291
        224.0.0.0        240.0.0.0            在链路上        26.26.26.1    257
        224.0.0.0        240.0.0.0            在链路上       192.168.8.1    291
        224.0.0.0        240.0.0.0            在链路上    169.254.254.97    281
  255.255.255.255  255.255.255.255            在链路上         127.0.0.1    331
  255.255.255.255  255.255.255.255            在链路上      192.168.18.1    281
  255.255.255.255  255.255.255.255            在链路上      192.168.0.26    311
  255.255.255.255  255.255.255.255            在链路上    169.254.75.113    291
  255.255.255.255  255.255.255.255            在链路上        26.26.26.1    257
  255.255.255.255  255.255.255.255            在链路上       192.168.8.1    291
  255.255.255.255  255.255.255.255            在链路上    169.254.254.97    281
===========================================================================
永久路由:
  网络地址          网络掩码  网关地址  跃点数
          0.0.0.0          0.0.0.0      192.168.8.2     默认
===========================================================================

Process finished with exit code 0

 */

//封装一个函数来执行命令
func execCommand(commandName string, params []string) bool {

	//执行命令
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)

	stdout, err := cmd.StdoutPipe()
	errReader, errr := cmd.StderrPipe()

	if errr != nil {
		fmt.Println("err:" + errr.Error())
	}

	//开启错误处理
	go handlerErr(errReader)

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()
	in := bufio.NewScanner(stdout)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		fmt.Println(cmdRe)
	}
	cmd.Wait()
	return true
}

//开启一个协程来错误
func handlerErr(errReader io.ReadCloser) {
	in := bufio.NewScanner(errReader)
	for in.Scan() {
		cmdRe := ConvertByte2String(in.Bytes(), "GB18030")
		fmt.Errorf(cmdRe)
	}
}

//对字符进行转码
func ConvertByte2String(byte []byte, charset Charset) string {
	var str string
	switch charset {
	case GB18030:
		var decodeBytes, _ = simplifiedchinese.GB18030.NewDecoder().Bytes(byte)
		str = string(decodeBytes)
	case UTF8:
		fallthrough
	default:
		str = string(byte)
	}
	return str
}
