package mycrypt

import (
	"reflect"
	"testing"
)

// Testene forutsetter et alfabet, som er definert i mycrypt.go som ALF_SEM03

func TestKrypter(t *testing.T) {
	type test struct {
		inputMessage []rune
		chiffer      int
		want         []rune
		errMsg       string
	}
	tests := []test{
		{inputMessage: []rune("w"), chiffer: 4, want: []rune("æ"), errMsg: ""},
		{inputMessage: []rune("0"), chiffer: 4, want: []rune("4"), errMsg: ""},
		//{inputMessage: []rune("Kjevik;SN39040;18.03.2022 01:50;6"), chiffer: 4, want: []rune("dnizmocdd7;484c5: 47 6466d45b94c."), errMsg: ""},
		//{inputMessage: []rune("dnizmocdd7;484c5: 47 6466d45b94c."), chiffer: len(ALF_SEM03) - 4, want: []rune("Kjevik;SN39040;18.03.2022 01:50;6"), errMsg: ""},
		{inputMessage: []rune("Kjevik;SN39040;18.03.2022 01:50;6"), chiffer: 4, want: []rune("bnizmoNcd7;484N5: 47 6466a45S94N."), errMsg: ""},
		{inputMessage: []rune("bnizmoNcd7;484N5: 47 6466a45S94N."), chiffer: len(ALF_SEM03) - 4, want: []rune("Kjevik;SN39040;18.03.2022 01:50;6"), errMsg: ""},
	}

	for _, tc := range tests {
		got, err := Krypter(tc.inputMessage, tc.chiffer)
		if !reflect.DeepEqual(got, tc.want) {
			t.Errorf("Feil, for chiffer %d, fikk %q, ønsket %q.", tc.chiffer, got, tc.want)
		}
		if (err == nil && tc.errMsg != "") || (err != nil && err.Error() != tc.errMsg) {
			t.Errorf("Feil, for chiffer %d, fikk feilmelding %q, ønsket %q.", tc.chiffer, err, tc.errMsg)
		}
	}
}

func TestSokIAlfabetet(t *testing.T) {
	type test struct {
		input rune
		want  int
	}
	tests := []test{
		{input: 'æ', want: 26},
		{input: 'a', want: 0},
	}

	for _, tc := range tests {
		got := sokIAlfabetet(tc.input, ALF_SEM03)
		if got != tc.want {
			t.Errorf("Feil, fikk %d, ønsket %d.", got, tc.want)
		}
	}
}
