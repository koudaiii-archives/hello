package models

import "strconv"

// User data
type User struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func NewUser(name string, age int) *User {
	return &User{
		Name: name,
		Age:  age,
	}
}

func (u *User) String() string {
	return u.Name + "(" + strconv.Itoa(u.Age) + ")"
}
