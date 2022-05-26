package main

import "fmt"

type User struct {
	firstname string
	surname   string
	age       int
	email     string
	password  string
}

func (u *User) editName(fname, sname string) {
	u.firstname = fname
	u.surname = sname
}

func (u *User) editAge(age int) {
	u.age = age
}

func (u *User) editEmail(email string) {
	u.email = email
}

func (u *User) editPass(pass string) {
	u.password = pass
}

func main() {
	user1 := User{
		"Lucas",
		"Bulhoes",
		23,
		"lucas.bulhoes@mercadolivre.com",
		"secret",
	}

	user1.editAge(25)

	fmt.Println(user1)
}
