package main

import (
	"testing"
)

func TestWallet(t *testing.T) {

	assertBalance := func(t testing.TB, wallet Wallet, want Bitcoin) {
    t.Helper()
		got := wallet.Balance()
		if got != want {
			t.Errorf("got %s want %s", got, want)
		}
	}

  assertError := func(t testing.TB, err error, want string) {
    t.Helper()
    if err == nil {
      t.Errorf("wanted an error but didn't get one")
    }

    if err.Error() != want {
      t.Errorf("got %q, want %q", err.Error(), want)
    }
  }

  assertNoError := func(t testing.TB, err error) {
    t.Helper()
    if err == nil {
      t.Errorf("wanted an error but didn't get one")
    }
  }

	t.Run("Deposit", func(t *testing.T) {
		wallet := Wallet{}

		wallet.Deposit(10)
		assertBalance(t, wallet, Bitcoin(10))

	})

	t.Run("Withdraw", func(t *testing.T) {
		wallet := Wallet{
			balance: 20,
		}

    err := wallet.Withdraw(10)

    assertNoError(t, err)
		assertBalance(t, wallet, Bitcoin(10))
	})

	t.Run("Withdraw insufficient funds", func(t *testing.T) {
    var startingBalance Bitcoin = Bitcoin(1)
		wallet := Wallet{
			balance: startingBalance,
		}

    err := wallet.Withdraw(10)

    assertError(t, err, ErrInsufficientFunds.Error())
		assertBalance(t, wallet, startingBalance)

	})
}
