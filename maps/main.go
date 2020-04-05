package main

import "fmt"

func main() {

	// var colors map[string]string
	colors := make(map[string]string) //above eqvuvalent

	// colors := map[string]string{
	// 	"red":   "#123456",
	// 	"green": "#111111",
	// }

	colors["White"] = "#fffff"
	fmt.Print(colors)
}
