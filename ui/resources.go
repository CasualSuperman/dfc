package ui

// Fonts
// DejaVu Sans Mono has good support for the characters we'll likely be using
// to build rooms.
// Monofur looks cool and also has good support, while the character shape
// looks like it belongs in a fantasy game; however, the transitions from thick
// to thin walls don't line up.

// ============ Termbox =============
// https://github.com/nsf/termbox-go/
// Pros:
//	Supports unicode
//	Cross-platform
// Cons:
//	In beta
//	Some problems with weird key combos
//	Not 256 Colors, only 8
// Notes:
//	Event-driven
