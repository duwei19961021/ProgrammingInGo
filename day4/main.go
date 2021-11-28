package main

import (
	"fmt"
	"net/http"
	"reflect"
	"sync"
)

// 反射

func reflectType(x interface{}) {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %s\n", v.Name())
}

type student struct {
	Name  string `ini:"name"`
	Score int    `json:"score"`
}

func main01() {
	stu1 := student{
		Name:  "zs",
		Score: 100,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(), t.Kind())

	// 遍历结构体字段信息
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		fmt.Printf("name:%s, index:%d, type:%v, json tag:%v\n", field.Name, field.Index, field.Type, field.Tag.Get("ini"))
	}
}

func SayHello(i int) {
	defer wg.Done()
	fmt.Printf("hello [%d]\n", i)
}

func curl() {
	client := http.Client{}
	req, _ := http.NewRequest("GET", "http://www.baidu.com", nil)
	respon, _ := client.Do(req)
	fmt.Printf("%s\n", respon.Status)
	respon.Body.Close()
	wg.Done()
}

var wg sync.WaitGroup

func main() {
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go curl()
	}
	wg.Wait()
	fmt.Println("main")
}
