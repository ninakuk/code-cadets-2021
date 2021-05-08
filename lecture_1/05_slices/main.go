package main

import "log"

func print(s []int) {
	log.Printf("elements: %v, len=%v, cap=%v\n", s, len(s), cap(s))
}

func main() {
	array := [5]int{1, 2, 3, 4, 5}//staticki array velicine 5 nije istog tipa kao isti velicine 4 npr
	print(array[:])

	slice := array[0:3]//od 0 do tri, kapacitet i dalje 5
	print(slice)

	slice = append(slice, 6)//vraca pointer na array, dodao 6 u SLICE
	print(slice)
	print(array[:])//umjesto 4 sad je 6

	emptySlice := make([]int, 5)//len i cap je 5
	print(emptySlice)

	emptySlice = append(emptySlice, 1)//uzmi array emptySlice, dodaj 1 na kraj i appendaj cap, odnosno poduplaj pa je cap sada 10
	print(emptySlice)

	emptySliceWithExtraCap := make([]int, 5, 10)//treci par je kapacitet
	print(emptySliceWithExtraCap)

	//////////////////////////////////////////

	sliceToIterate := []int{50, 60, 70, 80}

	//for each petlja zapravo
	for idx, val := range sliceToIterate {
		log.Printf("%v: %v\n", idx, val)//vraca indeks elementa i vrijednost
	}

	for _, val := range sliceToIterate {
		log.Printf("%v\n", val)
	}

	for val := range sliceToIterate {
		log.Printf("%v\n", val)
	}//ne iterira po vrijednostima nego po indeksima zapravo ! jer prva var je indeks, bez obzira kako je nazoves prva je indeks druga vrijednost, zato _,val
}
