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
	drawTypedBoundingBox(tb.ColorBlack, tb.ColorWhite)(x/3, 2*x/3, y/3, 2*y/3)
	fillInBox(x/3 + 1, 2*x/3 -1, y/3+1, 2*y/3-1, tb.ColorBlack, tb.ColorWhite)

	tb.Flush()
	time.Sleep(3 * time.Second)
}

func setDefaultCell(x, y int, r rune) {
	tb.SetCell(x, y, r, tb.ColorDefault, tb.ColorDefault)
}

func createSetKindOfCell(fg, bg tb.Attribute) func (int, int, rune) {
	return func(x, y int, r rune) {
		tb.SetCell(x, y, r, fg, bg)
	}
}

func fillInBox(xStart, xMax, yStart, yMax int, fg, bg tb.Attribute) {
	for i := xStart; i <= xMax; i++ {
		for j := yStart; j <= yMax; j++ {
			tb.SetCell(i, j, ' ', fg, bg)
		}
	}
}

func drawTypedBoundingBox(fg, bg tb.Attribute) func(int, int, int, int) {
	return func(xStart, xMax, yStart, yMax int) {
		drawBoundingBoxWithDrawFunc(xStart, xMax, yStart, yMax, createSetKindOfCell(fg, bg))
	}
}

func drawBoundingBoxWithDrawFunc(xStart, xMax, yStart, yMax int, draw func(int, int, rune)) {
	for i := xStart+1; i < xMax; i++ {
		draw(i, yStart, edges[0])
		draw(i, yMax,   edges[0])
	}
	for i := yStart+1; i < yMax; i++ {
		draw(xStart, i, edges[1])
		draw(xMax,   i, edges[1])
	}

	draw(xStart, yStart, corners[0])
	draw(xMax,   yStart, corners[1])
	draw(xStart, yMax,   corners[2])
	draw(xMax,   yMax,   corners[3])
}

func drawBoundingBox(xStart, xMax, yStart, yMax int) {
	drawBoundingBoxWithDrawFunc(xStart, xMax, yStart, yMax, setDefaultCell)
}
