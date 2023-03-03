package main

import (
	"errors"
	"fmt"
	"log"
	"os"
)

type Student struct {
	Name, Address, Job, Reason string
}

var students []Student = []Student{
	{
		Name:    "Nanda",
		Address: "Balikpapan",
		Job:     "Software Engineer",
		Reason:  "Sangat tertarik dengan Golang",
	},
	{
		Name:    "Windah",
		Address: "Jakarta",
		Job:     "Youtuber",
		Reason:  "Buat bikin konten Youtube",
	},
	{
		Name:    "Steve",
		Address: "Papua Barat",
		Job:     "Teknisi Laptop",
		Reason:  "Karena Golang populer",
	},
	{
		Name:    "Jamal",
		Address: "Makassar",
		Job:     "Designer",
		Reason:  "Ingin membuat program dengan Golang",
	},
}

func FindStudent(studentName string) (Student, error) {
	for _, value := range students {
		if value.Name == studentName {
			return value, nil
		}
	}

	return Student{}, errors.New("Student not found")
}

func main() {
	// get arguments
	inputs := os.Args

	// check length arguments
	if len(inputs) == 2 {
		result, err := FindStudent(inputs[1])
	
		if err == nil {
			fmt.Println("Nama :", result.Name)
			fmt.Println("Alamat :", result.Address)
			fmt.Println("Job :", result.Job)
			fmt.Println("Alasan mengikuti kelas Golang :", result.Reason)
		} else {
			log.Fatalln(err.Error())
		}
	} else {
		fmt.Println("Invalid arguments!")
	}
}
