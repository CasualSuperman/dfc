package main

import tb "github.com/nsf/termbox-go"
import "time"

var edges = []rune(`━┃`)
var corners = []rune(`┏┓┗┛`)
var intersections = []rune(`┳┫┻┣╋`)

func main() {
	if err := tb.Init(); err != nil {
		panic(err)
	} else {
		defer tb.Close()
		tb.HideCursor()
	}

	x, y := tb.Size()
	drawBoundingBox(0, x-1, 0, y-1)
	drawBoundingBox(x/3, 2*x/3, y/3, 2*y/3)

	tb.Flush()
	time.Sleep(3 * time.Second)
}

func setDefaultCell(x, y int, r rune) {
	tb.SetCell(x, y, r, tb.ColorDefault, tb.ColorDefault)
}

func drawBoundingBox(xStart, xMax, yStart, yMax int) {
	for i := xStart+1; i < xMax; i++ {
		setDefaultCell(i, yStart, edges[0])
		setDefaultCell(i, yMax,   edges[0])
	}
	for i := yStart+1; i < yMax; i++ {
		setDefaultCell(xStart, i, edges[1])
		setDefaultCell(xMax,   i, edges[1])
	}

	setDefaultCell(xStart, yStart, corners[0])
	setDefaultCell(xMax,   yStart, corners[1])
	setDefaultCell(xStart, yMax,   corners[2])
	setDefaultCell(xMax,   yMax,   corners[3])
}
