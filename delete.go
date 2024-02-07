package functional

import (
	"dima/read"
	"fmt"
	"io"
	"log"
	"os"
)

func DeleteUser() {
	id, check := read.ReadId()
	if !check {
		fmt.Println("incorrect id, try again")
		DeleteUser()
	}
	numUsr := findUser(id)
	if numUsr == -1 {
		fmt.Println("Usr with this id doesn't exist")
		return
	}
	file, err := os.OpenFile("users.txt", os.O_RDWR, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, byt)
	file.Seek(int64(numUsr*byt), 0)
	var newContent []byte
	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		newContent = append(newContent, buf...)
	}
	file.Seek(int64((numUsr-1)*byt), 0)
	file.Write(newContent)
	fmt.Println(int64(numUsr-1)*byt+int64(len(newContent)), "len delete")
	err = os.Truncate("users.txt", int64(numUsr-1)*byt+int64(len(newContent)))
	if err != nil {
		panic(err)
	}
}

