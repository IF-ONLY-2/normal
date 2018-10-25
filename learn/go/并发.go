package main

import (
	"fmt"
	"sync"
)

func main() {
	c1 := make(chan bool, 1)
	c2 := make(chan bool, 1)
	c1 <- true
	//a := <- c1
	//fmt.Println(a)
	wg := sync.WaitGroup{}
	wg.Add(2)
	go func() {
		defer wg.Done()
		for i := 1; i <= 26; i+=2 {
			<- c1
			fmt.Print(i, i+1)
			c2 <- true
		}
	}()

	go func() {
		defer wg.Done()
		for cc := 97; cc <=122; cc += 2 {
			<- c2
			fmt.Print(string(cc),string(cc+1))
			c1<-true
		}
	}()

	wg.Wait()
}
