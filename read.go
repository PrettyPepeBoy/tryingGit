package read

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

func ReadStdin() (string, error) {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		return "", err
	}
	input = strings.TrimSpace(input)
	return input, nil
}

func ReadButton() int {
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	i, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	return i
}

func ReadId() (uint32, bool) {
	fmt.Println("input user's id")
	id, err := ReadStdin()
	if err != nil {
		fmt.Println("incorrect input")
		return 0, false
	}
	identificator, err := strconv.Atoi(id)
	if err != nil {
		panic(err)
	}
	return uint32(identificator), true
}

func CheckFloat(input string) bool {
	_, err := strconv.ParseFloat(input, 64)
	return err == nil
}

// возвращаем слайс байтов для записи в файл
func CheckString(input string) ([]byte, bool) {
	placeHolder := []byte(input)
	for _, a := range placeHolder {
		if (a < uint8(65) || a > uint8(90)) && (a < uint8(97) || a > uint8(122)) {
			return nil, false
		}
	}
	var arr [20]byte
	for i, _ := range placeHolder {
		arr[i] = placeHolder[i]
	}
	return arr[:20], true
}

func AddStr(input string, err error) ([]byte, bool) {
	if err != nil {
		log.Fatal(err)
	}

	if byt, ok := CheckString(input); ok {
		return byt, true
	}
	return nil, false
}

func AddFloat(input string, err error) bool {
	if err != nil {
		log.Fatal(err)
	}
	if CheckFloat(input) {
		return true
	}
	return false
}

func addZero(input string) string {
	placeHolder := []byte(input)
	number, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}
	if placeHolder[0] != 48 && number < 10 {
		return "0" + input
	}
	return input
}

func CorrectData() (time.Time, bool) {
	var data string
	fmt.Println("please, input the year of your birth")
	input, err := ReadStdin()
	if err != nil {
		log.Fatal(err)
	}
	input = addZero(input)
	data = data + input

	fmt.Println("please, input the month of your birth")
	input, err = ReadStdin()
	if err != nil {
		log.Fatal(err)
	}
	input = addZero(input)
	data = data + input

	fmt.Println("please, input the day of your birth")
	input, err = ReadStdin()
	if err != nil {
		log.Fatal(err)
	}
	input = addZero(input)
	data = data + input

	layout := "20060102"
	t, err := time.Parse(layout, data)
	if err != nil {
		return t, false
	}
	return t, true
}
