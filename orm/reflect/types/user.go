package types

import "fmt"

type User struct {
	Name string
	age  int
}

func NewUser(name string, age int) User {
	return User{
		Name: name,
		age:  age,
	}
}

func NewUserPtr(name string, age int) *User {
	return &User{
		Name: name,
		age:  age,
	}
}

func (u User) GetAge() int {
	return u.age
}

func (u *User) ChangeName(name string) {
	u.Name = name
}

func (u User) private() {
	fmt.Println("private")
}
