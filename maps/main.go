package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#123456",
		"green": "#111111",
		"White": "#fffff",
	}

	for key, value := range colors {
		fmt.Println(key, value)
	}
}
