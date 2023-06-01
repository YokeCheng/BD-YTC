package main

import (
	"errors"
	"fmt"
)

func findUser(users []user, name string) (v *user, err error) {
	for _, u := range users {
		if u.Name == name {
			return &u, nil
		}
	}
	return nil, errors.New("not found")
}

func main() {
	u, err := findUser([]user{{"wang", "1024"}}, "wang")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(u.Name)

	if u, err := findUser([]user{{"wang", "1024"}}, "li"); err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println(u.Name)
	}
}
