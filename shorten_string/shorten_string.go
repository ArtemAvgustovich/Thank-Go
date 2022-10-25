package main

import (
	"fmt"
)

func main() {
	var text string
	var width int
	fmt.Scanf("%s %d", &text, &width)

	// Возьмите первые `width` байт строки `text`,
	// допишите в конце `...` и сохраните результат
	// в переменную `res`
	// ...
	var res string
	if len(text) > width {
		res = text[:width] + "..."
	} else {
		res = text
	}

	fmt.Println(res)
	fmt.Printf("%#v\n", res)
}
