package main

import "fmt"

func main() {
	var ch = make(chan int)
	for i := 0; i < 4; i++ {
		go func() {
			v, ok := <-ch
			if ok {
				if v == -1 {
					return
				}
				fmt.Println(v)
			} else {
				return
			}

		}()
	}

	ch <- 0
	ch <- 1
	ch <- -1

}
