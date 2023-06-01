package main

import (
	"fmt"
)

func main() {
	a := user{Name: "wang", Password: "1024"}
	b := user{"wang", "1024"}
	c := user{Name: "wang"}
	c.Password = "1024"
	var d user
	d.Name = "wang"
	d.Password = "1024"

	fmt.Println(a, b, c, d)
	fmt.Println(checkPassword(a, "haha"))
	fmt.Println(checkPassword2(&a, "haha"))
}

func checkPassword(u user, password string) bool {
	return u.Password == password
}

func checkPassword2(u *user, password string) bool {
	return u.Password == password
}
