package main

import "github.com/nsf/termbox-go"

type Input struct {
	endKey termbox.Key
	queue  chan termbox.Event
	ctrl   chan bool
}

func newInput() *Input {
	input := Input{
		endKey: termbox.KeyCtrlC,
		queue:	make(chan termbox.Event),
		ctrl:   make(chan bool, 2),
	}

	return &input
}

func (input *Input) start() {
	go poll(input)
}

func (input *Input) stop() {
	input.ctrl <- true
}

func poll(input *Input) {
	loop:
		for {
			select {
			case <- input.ctrl:
				break loop
			default:
				input.queue <- termbox.PollEvent()
			}
		}
}
