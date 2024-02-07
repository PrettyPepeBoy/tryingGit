package main

import (
	"dima/functional"
	"dima/read"
	"dima/write"
	"fmt"
	"os"
	"os/signal"
)

func main() {
	go func() {
		for {
			whatToDo()
		}
	}()
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)
	<-sigC
}

func whatToDo() {
	fmt.Println("Please input what you wanna do")
	fmt.Printf("Type |%d| to find user, |%d| to create new user, |%d| to sort users by name, |%d| to sort users by data,|%d| to delete user\n", 1, 2, 3, 4, 5)
	sigC := make(chan os.Signal, 1)
	signal.Notify(sigC, os.Interrupt)
	i := read.ReadButton()
	switch i {
	case 1:
		usr, ok := functional.CheckUsr()
		if !ok {
			fmt.Println("there is no such user")
		} else {
			fmt.Println(usr)
		}
	case 2:
		b, check := functional.CreateUser()
		if check {
			write.WriteInFile(b)
		}
	case 3:
		functional.SortByName()
	case 4:
		functional.SortByData()
	case 5:
		functional.DeleteUser()
	default:
		<-sigC
		fmt.Println("exit")
	}
}
