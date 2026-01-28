package main

import (
	"fmt"
	"sync"
	"time"
)

// go => GoRoutine
// GoRoutine => es un hilo de ejecución ligero virtual
// Channel => para comunicar cosas entre GoRoutine

func decirHola(canal chan<- string) {
	time.Sleep(1 * time.Second)
	canal <- "Hola desde la GoRoutine"
}

func imprimirMensaje(canal <-chan string) {
	msg := <-canal
	fmt.Println(msg)
}

func concurrency() {
	canal := make(chan string)
	go decirHola(canal)
	imprimirMensaje(canal)

	canal2 := make(chan int)
	go func() {
		for i := 0; i < 5; i++ {
			canal2 <- i
		}
		close(canal2)
	}()

	for num := range canal2 {
		fmt.Println("Número recibido: ", num)
	}

	//*****************
	//* Mutex es para reservar recursos
	//*****************

	// Ejemplo Mutex
	var contador int
	var mu sync.RWMutex

	// Writer (escritor)
	go func() {
		for i := 0; i < 5; i++ {
			mu.Lock()
			contador++
			mu.Unlock()
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// Reader (lector)
	for i := 0; i < 3; i++ {
		go func(id int) {
			for j := 0; j < 5; j++ {
				mu.RLock()
				fmt.Printf("Leyendo desde la GoRoutine %d: %d\n", id, contador)
				mu.RUnlock()
				time.Sleep(100 * time.Millisecond)
			}
		}(i)
	}

	time.Sleep(2 * time.Second)

}
