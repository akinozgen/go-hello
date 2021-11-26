package models

type Insan struct {
	id   int
	name string
	age  int
}

func NewInsan(id int, name string, age int) Insan {
	return Insan{
		id:   id,
		name: name,
		age:  age,
	}
}

func (i Insan) GetId() int {
	return i.id
}

func (i Insan) GetName() string {
	return i.name
}

func (i Insan) GetAge() int {
	return i.age
}

func (i *Insan) SetId(id int) *Insan {
	i.id = id
	return i
}

func (i *Insan) SetName(name string) *Insan {
	i.name = name
	return i
}

func (i *Insan) SetAge(age int) *Insan {
	i.age = age
	return i
}
