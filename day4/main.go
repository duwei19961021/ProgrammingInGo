package main

import (
	"fmt"
	"reflect"
)

// 反射

func reflectType(x interface{})  {
	v := reflect.TypeOf(x)
	fmt.Printf("type: %s\n",v.Name())
}

type student struct {
	Name string `ini:"name"`
	Score int `json:"score"`
}

func main() {
	stu1 := student{
		Name:  "zs",
		Score: 100,
	}

	t := reflect.TypeOf(stu1)
	fmt.Println(t.Name(),t.Kind())

	// 遍历结构体字段信息
	for i:=0;i<t.NumField();i++{
		field := t.Field(i)
		fmt.Printf("name:%s, index:%d, type:%v, json tag:%v\n",field.Name,field.Index,field.Type,field.Tag.Get("ini"))
	}
}