package main

import "fmt"

/*
import (
	"encoding/json"
	"fmt"
)

type animal struct{
	Name string `json:"name"`
}

func (a animal) Move(){
	fmt.Println("animal can move !")
}

type dog struct{
	Feet uint8 `json:"feet"`
	animal
}

func (d dog) Wang(){
	fmt.Println("dog can wang !")
}

func main(){
	d := new(dog)
	d.Feet = 4
	d.Name = "dog"

	b,_ := json.Marshal(d)
	fmt.Println(string(b))
}
*/

type animal interface{
	Speak()
}

type cat struct{}

func (c cat) Speak(){
	fmt.Println("喵喵喵！")
}

type dog struct{}

func (d dog) Speak(){
	fmt.Println("汪汪汪！")
}

func hit(x animal){
	x.Speak()
}

func main(){

	c := new(cat)

	hit(c)

	d := new(dog)

	hit(d)
}