package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	var c chan string = make(chan string)

	wg.Add(2)

	go Ping(c)
	go imprimir(c)

	wg.Wait()
	//var entrada string
	//fmt.Scanln(&entrada) - Desta forma não seria necessário o WaitGroups

}

var wg sync.WaitGroup

func Ping(c chan string) {
	for i := 0; i < 10; i++ { //Também poderia ser infinito, bastaria utilizar "ctrl + c" para sair
		c <- "Ping"
		c <- "Pong"
	}
	wg.Done()
}

func imprimir(c chan string) {
	for i := 0; i < 20; i++ {
		//msg := <-c
		//fmt.Println (msg) - Seria outra forma de printar, utilizando o Scanln, conforme comentário acima sobre WaitGroups.
		fmt.Println(<-c)
		time.Sleep(time.Second * 1)
	}
	wg.Done()
}
