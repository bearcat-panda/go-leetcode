package 面试题

import (
	"fmt"
	"testing"
)

var ch = make(chan int, 10)

func TestSerial(t *testing.T) {
	go func() {
		for {
			v, ok := <-ch
			if ok {
				fmt.Println(v)
			}
		}
	}()

	var flag = make(chan int, 1)
	for i := 0; i < 100; i++ {
		ch <- i

		if i == 99 {
			flag <- 1
		}
	}

	<-flag

}

func go1() {
	v, ok := <-chOne
	if ok {
		fmt.Println(v)
		go2()
	}
}
func go2() {
	v, ok := <-chTwo
	if ok {
		fmt.Println(v)
		go3()
	}
}
func go3() {
	v, ok := <-chThree
	if ok {
		fmt.Println(v)
		go1()
	}
}

var chOne = make(chan int, 1)
var chTwo = make(chan int, 1)
var chThree = make(chan int, 1)

func TestChannel(t *testing.T) {
	go go1()
	for i := 0; i < 100; i++ {
		switch i % 3 {
		case 0:
			chOne <- i
		case 1:
			chTwo <- i
		case 2:
			chThree <- i
		}
	}
}
