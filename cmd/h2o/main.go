package main

import (
	"fmt"
	"github.com/citrusinesis/semaphore-example/pkg/semaphore"
	"time"
)

var (
	sMutex     = semaphore.NewSemaphore(1)
	hydroQueue = semaphore.NewSemaphore(0)
	oxyQueue   = semaphore.NewSemaphore(0)
	barrier    = semaphore.NewBarrier(3)

	bond = bondInitialize(0)

	hydrogen uint = 0
	oxygen   uint = 0
)

func genHydrogen() {
	for {
		sMutex.Wait()
		hydrogen++

		if hydrogen >= 2 && oxygen >= 1 {
			hydroQueue.Signal()
			hydroQueue.Signal()
			hydrogen -= 2

			oxyQueue.Signal()
			oxygen--
		} else {
			sMutex.Signal()
		}

		hydroQueue.Wait()
		fmt.Println("Hydrogen Bond")

		bond()
		barrier.BarrierWait()
	}
}

func genOxygen() {
	for {
		sMutex.Wait()
		oxygen++

		if hydrogen >= 2 {
			hydroQueue.Signal()
			hydroQueue.Signal()
			hydrogen -= 2

			oxyQueue.Signal()
			oxygen--
		} else {
			sMutex.Signal()
		}

		oxyQueue.Wait()
		fmt.Println("Oxygen Bond")

		bond()
		barrier.BarrierWait()
		sMutex.Signal()
	}
}

func bondInitialize(i int) func() {
	return func() {
		i++
		if i%3 == 0 {
			fmt.Printf("** Molecule no. %d created**\n\n", i/3)
		}
		time.Sleep(time.Second * 1)
	}
}

func main() {
	go genHydrogen()
	go genHydrogen()
	go genOxygen()

	select {}
}
