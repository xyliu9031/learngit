package model

type person struct {
	Name string //保外可以使用p.Name
	age  int    //包外不可以使用p.age
}

func NewPerson(name string) *person {
	return &person{
		Name: name,
	}
}

func (p *person) SetAge(age int) {

}
