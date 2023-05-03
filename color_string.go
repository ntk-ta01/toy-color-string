package toy_color_string

import (
	"fmt"
	"strings"
)

const escape = "\x1b"

type ColorString struct {
	color    Color
	graphics []Graphic
	str      string
}

func NewColorString(color Color, str string) *ColorString {
	c := new(ColorString)
	c.color = color
	c.str = str
	return c
}

func (c *ColorString) AddGraphic(graphic Graphic) {
	c.graphics = append(c.graphics, graphic)
}

func (c *ColorString) arguments() string {
	arguments := make([]string, len(c.graphics)+1)
	arguments[0] = string(c.color)
	for i, graphic := range c.graphics {
		arguments[i+1] = string(graphic)
	}
	return strings.Join(arguments, ";")
}

func (c *ColorString) String() string {
	return fmt.Sprintf("%s[%sm%s%s[%sm", escape, c.arguments(), c.str, escape, RESET)
}
