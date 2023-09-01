package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	person "tugas1/model"
)

//

func json_Callback() {
	if len(os.Args) != 2 {
		fmt.Println("Usage: go run biodata.go <Student_code>")
		return
	}
	inputCode := os.Args[1]

	jsonFile := filepath.Join("helper", "participants.json")

	data, err := os.ReadFile(jsonFile)
	if err != nil {
		log.Fatal(err)
	}
	var people person.People

	if err := json.Unmarshal(data, &people); err != nil {
		log.Fatal(err)
	}

	found := false
	for _, p := range people.People {
		if p.Student_code == inputCode {
			found = true

			fmt.Printf("ID: %s\n", p.Id)
			fmt.Printf("Kode Peserta: %s\n", p.Student_code)
			fmt.Printf("Nama Peserta: %s\n", p.Student_name)
			fmt.Printf("Alamat Peserta: %s\n", p.Student_address)
			fmt.Printf("Pekerjaan Peserta: %s\n", p.Student_occupation)
			fmt.Printf("Alasan Peserta: %s\n", p.Joining_reason)

			break
		}
	}

	if !found {
		fmt.Println("Peserta dengan kode", inputCode, "tidak ditemukan dalam data")
	}
}

func main() {
	json_Callback()
}