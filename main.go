package main

import (
	"fmt"
	"github.com/lionsoul2014/ip2region/binding/golang/ip2region"
)

func main() {
	fmt.Println("err")
	region, err := ip2region.New("../../data/ip2region.db")
	defer region.Close()
	if err != nil {
		fmt.Println(err)
		return
	}
	ip, err := region.MemorySearch("47.102.212.119")
	fmt.Println(ip, err)
	ip, err = region.BinarySearch("47.102.212.119")
	fmt.Println(ip, err)
	ip, err = region.BtreeSearch("127.0.0.1")
	fmt.Println(ip, err)
}
