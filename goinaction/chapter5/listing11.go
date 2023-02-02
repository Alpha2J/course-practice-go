package main

import "fmt"

type user struct {
	name  string
	email string
}

func (u user) notify() {
	fmt.Println("sending user email to %s %s", u.name, u.email)
}

func (u *user) changeEmail(email string) {
	u.email = email
}

func main() {
	bill := user{
		"bill",
		"bill@gmail.com",
	}
	bill.notify()

	lisa := &user{"lisa", "lisa@gamil.com"}
	lisa.notify()

	bill.changeEmail("new-bill@gmail.com")
	bill.notify()
}
