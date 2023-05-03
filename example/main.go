package main

import (
	"fmt"
	"os"
	"strings"

	color_string "github.com/ntk-ta01/toy-color-string"
)

var (
	Output = os.Stdout
)

func main() {
	colorString := color_string.NewColorString(color_string.BLUE, "yeah")
	fmt.Println(colorString)
	colorString256 := color_string.NewColorString(color_string.To256Color(148), "let's go")
	colorString256.AddGraphic(color_string.ITALIC)
	fmt.Println(colorString256)
	rgb_colorString := color_string.NewColorString(color_string.ToRGBColor(255, 0, 0), "hello")
	fmt.Println(rgb_colorString)
	fmt.Println("world")
	fmt.Println()

	PASS := color_string.GREEN
	FAIL := color_string.RED
	RUN := color_string.BLUE
	message := `=== RUN   Test1
=== PAUSE Test1
=== RUN:  Test2
--- PASS: Test2 (0.02s)
=== CONT  Test3
=== RUN   Test3/A
=== RUN   Test3/B
=== RUN   Test3/C
--- FAIL: Test3
	--- PASS: Test3/A (0.00s)
	--- FAIL: Test3/B (0.00s)
	--- PASS: Test3/C (0.01s)`
	lines := strings.Split(message, "\n")
	for _, line := range lines {
		var color_line *color_string.ColorString
		if strings.HasPrefix(line, "--- PASS") {
			color_line = color_string.NewColorString(PASS, line)
		} else if strings.HasPrefix(line, "--- FAIL") {
			color_line = color_string.NewColorString(FAIL, line)
		} else if strings.HasPrefix(line, "--- FAIL") {
			color_line = color_string.NewColorString(FAIL, line)
		} else if strings.HasPrefix(line, "=== RUN") {
			color_line = color_string.NewColorString(RUN, line)
		} else {
			color_line = color_string.NewColorString(color_string.BLACK, line)
		}
		fmt.Println(color_line)
	}
}
