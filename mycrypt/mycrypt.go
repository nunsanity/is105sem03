package mycrypt

var ALF_SEM03 []rune = []rune("abcdefghijklmnopqrstuvwxyzæøå0123456789.,:; KSN")

func Krypter(melding []rune, alphabet []rune, chiffer int) []rune {
	kryptertMelding := make([]rune, len(melding))
	for i := 0; i < len(melding); i++ {
		indeks := SokIAlfabetet(melding[i], alphabet)
		if indeks+chiffer >= len(alphabet) {
			kryptertMelding[i] = alphabet[indeks+chiffer-len(alphabet)]
		} else {
			kryptertMelding[i] = alphabet[indeks+chiffer]
		}
	}
	return kryptertMelding
}

func Dekrypter(kryptertMelding []rune, alphabet []rune, chiffer int) []rune {
	dekryptertMelding := make([]rune, len(kryptertMelding))
	for i := 0; i < len(kryptertMelding); i++ {
		indeks := SokIAlfabetet(kryptertMelding[i], alphabet)
		if indeks-chiffer < 0 {
			dekryptertMelding[i] = alphabet[indeks-chiffer+len(alphabet)]
		} else {
			dekryptertMelding[i] = alphabet[indeks-chiffer]
		}
	}
	return dekryptertMelding
}

func SokIAlfabetet(symbol rune, alfabet []rune) int {
	for i := 0; i < len(alfabet); i++ {
		if symbol == alfabet[i] {
			return i
			break
		}
	}
	return -1
}
