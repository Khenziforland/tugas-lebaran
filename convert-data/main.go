package main

import (
	"fmt"
	"strconv"
	"strings"
)

type User struct {
	Name    string
	Age     int
	Address string
}

func ConvertData(data string) User {
	// Dapatkan Tiap Data
	splitted := strings.Split(data, ",")

	// Ubah tipe data age
    age, _ := strconv.Atoi(splitted[1])

    return User{
		Name: splitted[0], 
		Age: age, 
		Address: splitted[2]}
}

func main() {
	data := "Edo,20,Jakarta"
    user := ConvertData(data)
    fmt.Printf("%+v\n", user)

    data = "Budi,30,Bandung"
    user = ConvertData(data)
    fmt.Printf("%+v\n", user)
}
