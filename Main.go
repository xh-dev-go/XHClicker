package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	hook "github.com/robotn/gohook"
	"time"
)

func Click(left bool, done chan bool) {
	fmt.Printf("Start click: %v\n", left)
	ticker := time.NewTicker(50 * time.Millisecond)
	defer ticker.Stop()

	for {
		select {
		case <-done:
			fmt.Println("finish click event")
			return
		case <-ticker.C:
			if left {
				robotgo.Click("left", false)
				fmt.Println("left click")
			} else {
				robotgo.Click("right", false)
				fmt.Println("right click")
			}
		}
	}
}

var done chan bool = nil

func handle(event hook.Event) {
	//var done chan bool = nil
	if event.Button == 5 {
		fmt.Println("start right click")
		if done != nil {
			fmt.Println("Done is not nil")
			done <- true
		}
		done = make(chan bool)
		go Click(false, done)
	} else if event.Button == 4 {
		fmt.Println("start left click")
		if done != nil {
			fmt.Println("Done is not nil")
			done <- true
		}
		done = make(chan bool)
		go Click(true, done)
	} else if event.Button == 3 {
		fmt.Println("stop clicking")
		if done != nil {
			fmt.Println("Done is not nil")
			done <- true
			done = nil
		} else {
			fmt.Println("nothing happen")
		}
	}
}

func Add() {
	hook.Register(hook.KeyDown, []string{"q", "ctrl", "shift"}, func(e hook.Event) {
		fmt.Println("Done")
		hook.End()
	})

	hook.Register(hook.MouseDown, []string{}, func(event hook.Event) {
		handle(event)
	})

	s := hook.Start()
	<-hook.Process(s)
}

func main() {
	Add()
}
