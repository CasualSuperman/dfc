package main

import termbox "github.com/nsf/termbox-go"
import "fmt"

func main() {
	termbox.Init()
	defer termbox.Close()

	termbox.SetInputMode(termbox.InputEsc)

	termbox.Flush()

	for i := 0; i < 7; i++ {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			fmt.Println(ev)
		}
	}
}
