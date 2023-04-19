package main

import (
	"Macintosh_HD/Users/Nuns/is105/is105sem03/mycrypt"
	"fmt"
)

func main() {

	kryptertRune := mycrypt.Krypter([]rune("Kjevik;SN39040;18.03.2022 01:50;6"), mycrypt.ALF_SEM03, 4)
	kryptertString := string(kryptertRune)
	fmt.Println(kryptertString)

	alfLength := len(mycrypt.ALF_SEM03)

	dekryptertRune := mycrypt.Krypter(kryptertRune, mycrypt.ALF_SEM03, alfLength-4)
	dekryptertString := string(dekryptertRune)
	fmt.Println(dekryptertString)

}

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := mycrypt.SokIAlfabetet(melding[i], alphabet)
		kryptertMelding[i] = alphabet[(indeks+chiffer)%len(alphabet)]
	}
	return kryptertMelding
}
