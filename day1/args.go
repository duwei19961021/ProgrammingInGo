package main

import (
	"fmt"
	"github.com/pkg/errors"
	"log"
	"os"
	"strconv"
)

func AddToSlice(s *[]string) error {
	if cap(*s) == 100 {
		return errors.New("slice is empty!")
	}
	fmt.Println(cap(*s))
	return nil
}

func main1() {
	/*
	fmt.Println(os.Args)
	var str []string = make([]string,100)
	err := AddToSlice(&str)
	fmt.Println(err)
	*/

	if _,err:=os.Open("a.txt");err!=nil{
		log.Fatal(err)
	}
}

/*
 make函数返回该类型的引用（和cpp里的引用差不多，不需要解引用）
 */

type polar struct {
	radius float64
	shuai float64
}

type cartesian struct {
	x float64
	y float64
}

func main2() {
	var messages chan string = make(chan string, 10) // 创建容量为10的管道
	messages <- "Leader"
	var msg1 string = <-messages
	var msg2 string = <-messages
	fmt.Println(msg1,msg2)
}

func main() {
	var s string = "duwei1996杜"
	fmt.Println(len(s))
	ret:= strconv.Itoa(65)
	fmt.Println(ret+"1")
}