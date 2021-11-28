package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// goroutine

func test() {
	for i := 1; i < 10; i++ {
		fmt.Println("test: hello world" + strconv.Itoa(i))
		time.Sleep(time.Second * 1)
	}
}

func Compute(num int64) {
	var ret int64 = 1
	var i int64
	for i = 1; i <= num; i++ {
		ret *= num
	}
	lock.Lock()
	myMap[num] = ret
	lock.Unlock()
	wg.Done()
}

var (
	wg    sync.WaitGroup
	myMap = make(map[int64]int64, 10)
	lock  sync.Mutex
)

func main_goroutine() {
	/*
		go test()

		for i := 0 ;i < 10; i++ {
			fmt.Println("main: hello world" + strconv.Itoa(i))
			time.Sleep(time.Second * 1)
		}
	*/

	var i int64
	for i = 1; i <= 200; i++ {
		wg.Add(1)
		go Compute(i)
	}

	lock.Lock()
	for k, v := range myMap {
		fmt.Printf("map[%d]=%d\n", k, v)
	}
	lock.Unlock()

	wg.Wait()
}

////////////////////////////////// channel

type Cat struct {
	name string
	age  int
}

func main_chain() {
	/*
		var intChan chan int
		intChan = make(chan int,10)
		intChan <- 1
		intChan <- 2
		intChan <- 3

		num1 := <- intChan
		num2 := <- intChan
		num3 := <- intChan
		fmt.Println(num1,num2,num3)
	*/

	allChain := make(chan interface{}, 3)
	allChain <- 10
	allChain <- "tomcat"
	cat := Cat{
		name: "jack",
		age:  4,
	}
	allChain <- cat

	<-allChain
	<-allChain

	c := <-allChain
	fmt.Println(c.(Cat).name)
}

/////////////////////// chain & goroutine

func WriteData(intChain chan int) {
	for i := 1; i <= 50000; i++ {
		intChain <- i
		fmt.Printf("write data: %d\n", i)
	}
	close(intChain)
}

func ReadData(intChain chan int, exitChain chan bool) {
	for {
		v, ok := <-intChain
		if !ok {
			break
		}
		fmt.Printf("read data: %d\n", v)
	}
	exitChain <- true
	close(exitChain)
}

func main() {
	//dataChain := make(chan int, 50)
	exitChain := make(chan bool, 1)

	//go WriteData(dataChain)

	//go ReadData(dataChain, exitChain)

	for {
		_, ok := <-exitChain // chan在接收和发送会阻塞，阻塞条件是接收是缓冲空或者发送时缓冲满；
		if !ok {
			break
		}
	}
}
