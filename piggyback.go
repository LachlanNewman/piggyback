package main

import(
	"fmt"
	"github.com/LachlanNewman/user"
)

func main()  {
	u := user.NewUser("test","test")
	fmt.Println(u.GetFirstName())
	fmt.Println(u.GetLastName())
}