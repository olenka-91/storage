package main

import (
	"fmt"
	"log"

	"github.com/olenka-91/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	file, err := st.Upload("test1.txt", []byte("hello"))
	if err != nil {
		log.Fatal(err)
	}
	resFile, err := st.GetByID(file.ID)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(st.Files)
	fmt.Println("It is restored", resFile)
}
