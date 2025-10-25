package main

import (
	"fmt"
	// "../../demo 12/model"
)

type Animal struct {
	Age    int
	Weight float64
}

func (an *Animal) ShowInfo() {
	fmt.Printf("动物的年龄为%v,动物的体重为%v\n", an.Age, an.Weight)
}

type Cat struct {
	Animal
}

func main() {
	// P := model.NewPerson("liuxinyu")

	cat := &Cat{}
	cat.Animal.Age = 3
	cat.Age = 4 //两种方式都可以
	cat.Animal.Weight = 10.1
	// println(cat.Animal.Age, cat.Age)
	cat.Animal.ShowInfo()

	cat.Animal.ShowInfo()
	println(cat.Animal.Age, cat.Age)
}
