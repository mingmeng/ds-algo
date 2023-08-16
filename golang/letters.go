package golang

import (
	"fmt"
	"sync"
)

var stopCh chan int

func printLetters(letters byte, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for {

		select {
		case <-ch:
			fmt.Printf("letters: %s\n", string(letters))
			ch <- 1
		case <-stopCh:
			fmt.Println("stop output!")
			break
		}
	}
}

func letters() {
	flagCh := make(chan int)
	stopCh = make(chan int)

	wg := new(sync.WaitGroup)

	wg.Add(2)
	for i := 0; i < 26; i += 2 {
		go printLetters(byte('0'+i), flagCh, wg)
		go printLetters(byte('0'+i+1), flagCh, wg)
	}

	stopCh <- 1
}
