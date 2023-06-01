package main

import (
	"fmt"
	"github/BD-YTC/YTC-20230512/goprinciplespractices/structures"
)

type user = structures.User

func main() {
	a := user{Name: "wang", Password: "1024"}
	a.ResetPassword("2048")
	fmt.Println(a.CheckPassword("2048"))
}
