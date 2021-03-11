package main

import (
	"fmt"
	"time"
)

func main() {

	c := make(chan string) //send values over channels
	go count("sheep", c)

	msg := <-c
	fmt.Println(msg)

}

func count(thing string, c chan string) {
	for i := 1; true; i++ {
		//fmt.Println(i, thing)
		c <- thing //waits until receiver is ready to receive
		time.Sleep(time.Millisecond * 500)
	}
	close(c)
	//only close as sender, not receiver as sender may not know channel has been closed
}

//USE OF CHANNELS - When a goroutine requires a value from a channel it will wait until the value is available
//eg goroutine2 processing data to provide the value for goroutine1 to continue
