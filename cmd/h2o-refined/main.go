package main

import (
	"fmt"
	"github.com/citrusinesis/semaphore-example/pkg/semaphore"
)

var (
	hydrogen uint = 0

	mutex          = semaphore.NewBinarySemaphore()
	pairOfHydrogen = semaphore.NewSemaphore(0)
	oxygen         = semaphore.NewSemaphore(0)
)

func hReady() {
	for {
		mutex.Wait()
		hydrogen++
		fmt.Println("Hydrogen Created")

		if hydrogen%2 == 0 {
			pairOfHydrogen.Signal()
		}
		mutex.Signal()

		oxygen.Wait()
	}
}

func oReady() {
	for {
		fmt.Println("Oxygen Created")
		pairOfHydrogen.Wait()
		makeWater()
		oxygen.Signal()
		oxygen.Signal()
	}
}

func makeWater() {
	mutex.Wait()
	hydrogen -= 2
	mutex.Signal()

	fmt.Println("CREATED!")
	fmt.Printf("H: %d\n", hydrogen)
}

func main() {

	go hReady()
	go hReady()
	go oReady()

	select {}
}
