package toy_color_string

import (
	"fmt"
)

type Color string

const (
	BLACK          Color = "30"
	RED            Color = "31"
	GREEN          Color = "32"
	YELLOW         Color = "33"
	BLUE           Color = "34"
	MAGENTA        Color = "35"
	CYAN           Color = "36"
	WHITE          Color = "37"
	BRIGHT_BLACK   Color = "90"
	BRIGHT_RED     Color = "91"
	BRIGHT_GREEN   Color = "92"
	BRIGHT_YELLOW  Color = "93"
	BRIGHT_BLUE    Color = "94"
	BRIGHT_MAGENTA Color = "95"
	BRIGHT_CYAN    Color = "96"
	BRIGHT_WHITE   Color = "97"
)

func To256Color(id int) Color {
	if id < 0 || 255 < id {
		id = 0
	}
	return Color(fmt.Sprintf("38;5;%d", id))
}

func ToRGBColor(r int, g int, b int) Color {
	if r < 0 || 255 < r {
		r = 0
	}
	if g < 0 || 255 < g {
		g = 0
	}
	if b < 0 || 255 < b {
		b = 0
	}
	return Color(fmt.Sprintf("38;2;%d;%d;%d", r, g, b))
}
