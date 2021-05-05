package bank

import "testing"

func TestWallet(t *testing.T) {

	confirmBalance := func(t *testing.T, w Wallet, want Bitcoin) {
		t.Helper()
		got := w.Balance()

		if got != want {
			t.Errorf("got: '%s' want: '%s'", got, want)
		}
	}

	t.Run("check inicial balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(25.0))

		confirmBalance(t, wallet, Bitcoin(25.0))
	})

	t.Run("check method Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		wallet.Withdraw(5)

		confirmBalance(t, wallet, Bitcoin(15))
	})

}
