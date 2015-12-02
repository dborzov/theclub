package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pirate(name string, out chan string) {
	cnt := 0
	for {
		cnt++
		time.Sleep((time.Duration(rand.Intn(1300))) * time.Millisecond)
		if cnt > 3 {
			out <- fmt.Sprintf("%v is tired of this shit and leaves\n", name)
			return
		}
		res := fmt.Sprintf("%v says: heyo for %vth time! \n", name, cnt)
		out <- res
	}
}

func main() {
	log := make(chan string)
	for i := 65; i < 68; i++ {
		go pirate(fmt.Sprintf("Pirate %v", string(rune(i))), log)
	}

	go func() {
		for l := range log {
			fmt.Printf(l)
		}
	}()
	time.Sleep(10000 * time.Millisecond)
}
