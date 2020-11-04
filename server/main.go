package main

import (
	"encoding/json"
	"fmt"
	"log"
	"nat_proxy/server/lib/network"
	"nat_proxy/server/web"
	"nat_proxy/server/web/common"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	quit := make(chan os.Signal, 1)
	// ctrl+c->SIGINT, kill -9 -> SIGKILL
	signal.Notify(quit, syscall.SIGINT, syscall.SIGKILL)
	// start server
	// 1.web server
	go web.Web(common.WebAddress)
	go func() {
		publicIP, err := network.PublicIPInfo()
		if err != nil {
			log.Fatal(err)
		}
		b, err := json.Marshal(publicIP)
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(string(b))
	}()
	log.Printf("Servers killed by %v signal.", <-quit)

}
