package bank

import "testing"

func TestWallet(t *testing.T) {

	t.Run("verify inicial balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(25.0))

		got := wallet.Balance()
		want := Bitcoin(25.0)

		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	})

}
