package main

import "fmt"

func main() {
	colors := map[string]string{
		"red":   "#ff000",
		"green": "test",
	}
	var colors2 map[string]string
	colors2["hello"] = "test"

	colors3 := make(map[string]string)
	colors3["test"] = "test"

	delete(colors, "red")

	for key, value := range colors {
		fmt.Println(key, value)
	}
	fmt.Println(colors)
}
