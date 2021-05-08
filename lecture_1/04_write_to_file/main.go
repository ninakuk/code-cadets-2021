package main

import (
	"fmt"
	"log"
	"os"

	"github.com/pkg/errors"
)

func countToTen(f *os.File) {
	for i := 10; i >= 0; i-- {
		defer f.WriteString(fmt.Sprint(i) + "\n")//zapravo stavlja na stack, ide unazad, ako file vec postoji brise sadrzaj unutra i pise novi
	}
}

func main() {
	f, err := os.Create("output.txt")
	if err != nil {
		log.Fatal(
			errors.WithMessage(err, "opening a file"),
		)
	}

	//defer omogucuje da se cijeli main izvrsi prije nego se file zatvori
	defer f.Close()
	countToTen(f)
	f.Sync()
}
