package network

import (
	"fmt"
	"log"
	"testing"
)

func TestLocalIP(t *testing.T) {
	ip, err := LocalIP()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(ip)
}
