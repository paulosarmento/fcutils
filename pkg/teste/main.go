package main

import "fmt"

func main() {
	evento := []string{"test1", "test2", "test3", "test4"}
	// evento = evento[0:]
	evento = append(evento[:0], evento[1:]...)
	// fmt.Println(evento1)
	// 01234
	fmt.Println(evento)

}
