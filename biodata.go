package main

import (
	"fmt"
	"os"
	"strconv"
)

type User struct {
	NomorAbsen uint8
	Nama       string
	Alamat     string
	Pekerjaan  string
	Alasan     string
}

func main() {
	nomorAbsen, _ := strconv.Atoi(os.Args[1])

	result := checkUserExist(uint8(nomorAbsen))

	fmt.Println(result)
}

func checkUserExist(input uint8) User {
	var dataSearched User

	users := []User{
		{NomorAbsen: 1, Nama: "Jekkal", Alamat: "Meikarta", Pekerjaan: "Programmer", Alasan: "Bosan"},
		{NomorAbsen: 3, Nama: "Erick", Alamat: "Mediterranian", Pekerjaan: "Designer", Alasan: "Duit"},
		{NomorAbsen: 2, Nama: "Jamal", Alamat: "Yvse", Pekerjaan: "Business Intelligence", Alasan: "Lokasi"},
	}

	for _, value := range users {
		if value.NomorAbsen == input {
			dataSearched = value
			break
		}
	}
	return dataSearched
}
