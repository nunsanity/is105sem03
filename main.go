package main

import (
	"fmt"

	"github.com/nunsanity/is105sem03/mycrypt"
)

func main() {
	input := []rune("Kjevik;SN39040;18.03.2022 01:50;6")

	kryptertRune, err := mycrypt.Krypter(input, 4)
	if err != nil {
		fmt.Println("Krypter error:", err)
		return
	}
	kryptertString := string(kryptertRune)
	fmt.Println(kryptertString)

	alfLength := len(mycrypt.ALF_SEM03)

	dekryptertRune, err := mycrypt.Krypter(kryptertRune, alfLength-4)
	if err != nil {
		fmt.Println("Dekrypter error:", err)
		return
	}
	dekryptertString := string(dekryptertRune)
	fmt.Println(dekryptertString)
}
