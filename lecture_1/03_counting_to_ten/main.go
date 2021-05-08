package main

import "log"
import "fmt"

func main() {

	for i := 0; i <= 10; i++ {
		log.Println(i)
		//fmt.Sprintf()
	}

	//zapravo while petlja, go nema while ili do while, sve ide preko for
	i := 0
	for i <= 10 {
		log.Println(i)
		i++
	}

	i = 0
	for {
		fmt.Println(i)
		i++
		if i == 11 {
			break
		}
	}
}
