package main

import (
	"fmt"
	"math/rand"
	"time"
)

func pirate(name string) {
	cnt := 0
	for {
		cnt++
		time.Sleep((200 + time.Duration(rand.Intn(1300))) * time.Millisecond)
		fmt.Printf("%v says: heyo for %vth time! \n", name, cnt)
	}
}

func main() {
	for i := 65; i < 90; i++ {
		go pirate(fmt.Sprintf("Pirate%v", string(rune(i))))
	}
	time.Sleep(5000 * time.Millisecond)
}
