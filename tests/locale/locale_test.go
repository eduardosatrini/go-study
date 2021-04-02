package locale

import "testing"

func TestAdress(t *testing.T) {
	value := "street valley"
	esperado := "Street"
	recebido := Adress(value)

	if recebido != esperado {
		t.Error("Error")
	}
}
