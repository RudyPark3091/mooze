package ui

import (
	"fmt"
)

func LoadAlternateScreen() {
	fmt.Print("\033[?1049h\033H")
}

func UnloadAlternateScreen() {
	fmt.Print("\033[?1049l")
}

func Print(s string) {
	fmt.Print(s, "\r")
}

func Println(s string) {
	fmt.Println(s, "\r")
}

func KeyBindings() {
	main := Fg(
		"Mooze: Yet another REST api test tool for command-line users",
		MAGENTA)
	Println(main)

	Println("Request")
	Println("- " + Fg("u", RED) + "rl: ")
	Println("- " + Fg("m", RED) + "ethod: ")
	Println("- " + Fg("b", RED) + "ody: ")
	Println("- " + Fg("h", RED) + "eader: ")

	Println("Operations")
	Println("- " + Fg("r", RED) + "efresh screen")
	Println("- " + Fg("s", RED) + "end request")
	Println("- " + Fg("q", RED) + "uit")
}
