package main

import (
	"github.com/nsf/termbox-go"
)

const (
	EventKey EventType = iota
	EventResize
	EventMouse
	EventError
	EventInterrupt
	EventRaw
	EventNone
)

type (
	Attr      uint16
	Key       uint16
	Modifier  uint8
	EventType uint8
)

type Event struct {
	Type   EventType // The type of event
	Key    Key       // The key pressed, if any
	Ch     rune      // The character of the key, if any
	Mod    Modifier  // A keyboard modifier, if any
	Err    error     // Error, if any
	MouseX int       // Mouse X coordinate, if any
	MouseY int       // Mouse Y coordinate, if any
}

func convertEvent(tBoxEvent termbox.Event) Event {
	return Event{
		Type:   EventType(tBoxEvent.Type),
		Key:    Key(tBoxEvent.Key),
		Ch:     tBoxEvent.Ch,
		Mod:    Modifier(tBoxEvent.Mod),
		Err:    tBoxEvent.Err,
		MouseX: tBoxEvent.MouseX,
		MouseY: tBoxEvent.MouseY,
	}
}
