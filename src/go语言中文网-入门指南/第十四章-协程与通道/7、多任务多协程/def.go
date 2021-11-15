package main

import (
	"time"
)

type Task struct {
	a int
}

func main() {

	in := make(chan *Task)

	go Send(in)

	out := make(chan *Task)
	for i := 0; i < 10; i++ {
		go Worker(in, out)
	}
}

func Send(in chan *Task) {
	for i := 0; i < 50; i++ {
		in <- &Task{i}
	}
	close(in)
}

func Worker(in, out chan *Task) {
	for {
		t := <-in
		process(t)
		out <- t
	}
}

func process(task *Task) {
	time.Sleep(1e9)
}
