package main

import "fmt"

type Person struct {
	name string
	age int
	hobby [3]string
	msg map[string]string
}

// 构造函数
func newPerson() *Person  {
	p := new(Person)
	m := make(map[string]string)
	p.msg = m
	return p
}

func (p *Person) PrintPerson()  {
	fmt.Println(p.name,p.age,p.hobby,p.msg)
}

func main() {
	p := newPerson()
	p.msg["addr"] = "hz"
	p.name = "dw"
	p.age = 25
	p.hobby = [3]string{"go","c","cpp"}

	p2 := newPerson()
	p2.msg["addr"] = "hz"
	p2.name = "dw"
	p2.age = 25
	p2.hobby = [3]string{"go","c","cpp"}

	p2.PrintPerson()
}
