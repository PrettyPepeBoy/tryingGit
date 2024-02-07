package functional

import (
	"dima/read"
	"encoding/binary"
	"fmt"
	"math/rand"
)

const byt = 50

type User struct {
	Id      uint32
	Name    string
	Surname string
	Year    uint16
	Month   uint16
	Day     uint16
}

type Users []User

func CreateUser() ([]byte, bool) {
	userSlc := make([]byte, 0, byt)
	id := randInt(100000, 999999)
	userSlc = binary.BigEndian.AppendUint32(userSlc, id)

	fmt.Println("please, input your user name")
	input, err := read.ReadStdin()
	if err != nil {
		panic(err)
	}
	data := make([]byte, 20)
	for i := range input {
		data[i] = input[i]
	}
	userSlc = append(userSlc, data...)

	fmt.Println("please, input you user surname")
	input, err = read.ReadStdin()
	if err != nil {
		panic(err)
	}
	data = make([]byte, 20)
	for i := range input {
		data[i] = input[i]
	}
	userSlc = append(userSlc, data...)

	t, check := read.CorrectData()
	if check {
		userSlc = binary.BigEndian.AppendUint16(userSlc, uint16(t.Year()))
		userSlc = binary.BigEndian.AppendUint16(userSlc, uint16(t.Month()))
		userSlc = binary.BigEndian.AppendUint16(userSlc, uint16(t.Day()))
	}

	return userSlc, true
}

func GetUserName(name, surname string) {
	Usr := getData("users.txt")
	for i := 0; i < len(Usr); i++ {
		if Usr[i].Name == name && Usr[i].Surname == surname {
			fmt.Println(Usr[i])
		}
	}
}

func randInt(i, j int) uint32 {
	return uint32(i + rand.Intn(j-i))
}

