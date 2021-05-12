package bank

import "testing"

func TestWallet(t *testing.T) {

	t.Run("check inicial balance", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(25.0))

		confirmBalance(t, wallet, Bitcoin(25.0))
	})

	t.Run("check method Withdraw", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(Bitcoin(20))
		err := wallet.Withdraw(Bitcoin(5))

		confirmInexistentErr(t, err)
		confirmBalance(t, wallet, Bitcoin(15))
	})

	t.Run("check negative balance", func(t *testing.T) {
		inicialBalance := Bitcoin(100)
		wallet := Wallet{}
		wallet.Deposit(inicialBalance)
		err := wallet.Withdraw(Bitcoin(110))

		confirmBalance(t, wallet, inicialBalance)
		checkErr(t, err, ErrorInsufficientFunds)
	})

}

func confirmInexistentErr(t *testing.T, result error) {
	t.Helper()

	if result != nil {
		t.Fatal("unexpected error")
	}
}

func confirmBalance(t *testing.T, w Wallet, want Bitcoin) {
	t.Helper()
	got := w.Balance()

	if got != want {
		t.Errorf("got: '%s' want: '%s'", got, want)
	}
}

func checkErr(t *testing.T, result error, expected error) {
	t.Helper()
	if result == nil {
		t.Fatal("Expected an error but none occurred.")
	}

	if result != expected {
		t.Errorf("result: '%s', expected: '%s'", result, expected)
	}
}
