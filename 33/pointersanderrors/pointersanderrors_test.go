package pointersanderrors

import "testing"

func TestWallet(t *testing.T) {

	t.Run("deposit", func(t *testing.T) {
		wallet := Wallet{}
		wallet.Deposit(10.00)
		got := wallet.Balance()
		want := Bitcoin(10.00)
		assertBalance(t, got, want)
	})
	t.Run("withdraw no error with available balance", func(t *testing.T) {
		wallet := Wallet{balance: 20.00}
		err := wallet.Withdraw(10.00)
		got := wallet.Balance()
		want := Bitcoin(10.00)
		assertNoError(t, err)
		assertBalance(t, got, want)
	})
	t.Run("withdraw error and fail without available balance", func(t *testing.T) {
		wallet := Wallet{balance: 20.00}
		err := wallet.Withdraw(30.00)
		got := wallet.Balance()
		want := Bitcoin(20.00)
		assertError(t, err, ErrInsufficientFunds)
		assertBalance(t, got, want)
	})

}

func assertBalance(t testing.TB, got Bitcoin, want Bitcoin) {
	t.Helper()
	if got != want {
		t.Errorf("got: %s want: %s", got, want)
	}
}

func assertError(t testing.TB, gotErr error, wantErr error) {
	t.Helper()
	if gotErr == nil {
		t.Fatalf("should have received error but didn't")
	}
	if gotErr != wantErr {
		t.Errorf("got error: %q want error: %q", gotErr, ErrInsufficientFunds)
	}
}

func assertNoError(t testing.TB, err error) {
	t.Helper()
	if err != nil {
		t.Fatalf("should not have received error but received %q", err)
	}
}
