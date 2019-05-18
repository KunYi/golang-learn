package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

type workitem struct {
	id   int
	data string
}

func genWork(out chan<- workitem) {
	for i := 0; i < 1000; i++ {
		out <- workitem{id: i, data: string(rand.Intn(1000))}
	}
}

func pipeStage(out chan<- workitem, in <-chan workitem) {
	for {
		d := <-in // need copy object
		d.data = d.data + " + stage"
		out <- d
	}
}

func consume(in <-chan workitem) {
	for {
		d := <-in // need copy object
		log.Println(fmt.Sprintf("WorkItem id: %d, Data:%s", d.id, d.data))
	}
}

func main() {
	log.Println(fmt.Sprintf("%s", "Hello world!"))
	log.Println(fmt.Sprintf("%s", "Channel and Pipeline demo"))
	it := make(chan workitem)
	et := make(chan workitem)
	go genWork(it)
	go consume(et)
	go pipeStage(et, it)

	/*
	 should to use sync.WaitGroup
	 the code juse demo waiting all channel run out
	*/
	time.Sleep(100 * time.Millisecond) // waiting for kick-off of channel
loop:
	select {
	case <-et:
		time.Sleep(1 * time.Second)
		break loop
	default:
		break
	}

	log.Println(fmt.Sprintf("%s", "End of all work"))
}
