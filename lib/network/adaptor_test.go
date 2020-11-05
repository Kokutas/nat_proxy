package network

import (
	"encoding/json"
	"fmt"
	"log"
	"testing"
)

func Test_adaptors(t *testing.T) {
	adaptors, err := adaptors()
	if err != nil {
		log.Fatal(err)
	}
	for _, adaptor := range adaptors {
		data, err := json.Marshal(adaptor)
		if err != nil {
			continue
		}
		fmt.Printf("%s\n", data)
	}
}

func TestAdaptors(t *testing.T) {
	ip, err := LocalIP()
	if err != nil {
		log.Fatal(err)
	}
	adaptor, err := Adaptors(ip)
	if err != nil {
		log.Fatal(err)
	}
	data, err := json.Marshal(adaptor)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s------------\n", data)
}
