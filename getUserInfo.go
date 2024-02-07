package functional

import (
	"bufio"
	"dima/read"
	"encoding/binary"
	"fmt"
	"io"
	"log"
	"os"
)

func CheckUsr() (User, bool) {
	id, ok := read.ReadId()
	if !ok {
		fmt.Println("incorrect id, try again")
		CheckUsr()
	}
	if num := findUser(id); num != -1 {
		file, err := os.OpenFile("users.txt", os.O_RDONLY, 0644)
		if err != nil {
			panic(err)
		}
		file.Seek(int64((num-1)*byt), 0)
		buf := make([]byte, byt)
		file.Read(buf)
		usr := User{
			Id:      binary.BigEndian.Uint32(buf),
			Name:    string(buf[4:24]),
			Surname: string(buf[24:44]),
			Year:    binary.BigEndian.Uint16(buf[44:46]),
			Month:   binary.BigEndian.Uint16(buf[46:48]),
			Day:     binary.BigEndian.Uint16(buf[48:50]),
		}
		return usr, true
	}
	return User{}, false
}

func getData(fileName string) Users {
	var usr Users
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	reader := bufio.NewReader(file)
	buf := make([]byte, byt)
	for {
		_, err := reader.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		Usr := User{
			Id:      binary.BigEndian.Uint32(buf),
			Name:    string(buf[4:24]),
			Surname: string(buf[24:44]),
			Year:    binary.BigEndian.Uint16(buf[44:46]),
			Month:   binary.BigEndian.Uint16(buf[46:48]),
			Day:     binary.BigEndian.Uint16(buf[48:50]),
		}
		usr = append(usr, Usr)
	}
	return usr
}

func findUser(id uint32) int {
	file, err := os.OpenFile("users.txt", os.O_RDONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	buf := make([]byte, byt)
	count := 1
	for {
		_, err := file.Read(buf)
		if err != nil {
			if err != io.EOF {
				panic(err)
			}
			break
		}
		if binary.BigEndian.Uint32(buf) == id {
			return count
		}
		count++
	}
	return -1
}
