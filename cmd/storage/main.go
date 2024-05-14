package main

import (
	"fmt"

	"github.com/olenka-91/storage/internal/storage"
)

func main() {
	st := storage.NewStorage()
	fmt.Println("It works", st)
}
