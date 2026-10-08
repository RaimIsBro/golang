package day7

import (
	"errors"
	"testing"
)

func TestSwap(t *testing.T) {
	a, b := 1, 2
	Swap(&a, &b)
	if a != 2 || b != 1 {
		t.Errorf("после Swap a = %d, b = %d; want 2, 1", a, b)
	}
}

func TestResetCopyVsPointer(t *testing.T) {
	u := User{Name: "Раим", Age: 30}
	ResetAgeCopy(u)
	if u.Age != 30 {
		t.Errorf("ResetAgeCopy не должна менять оригинал, Age = %d", u.Age)
	}
	ResetAge(&u)
	if u.Age != 0 {
		t.Errorf("ResetAge должна обнулить Age через указатель, Age = %d", u.Age)
	}
	ResetAge(nil) // не должна паниковать
}

func TestAccount(t *testing.T) {
	acc := NewAccount("Раим")
	if acc == nil {
		t.Fatal("NewAccount вернула nil")
	}
	if acc.Owner != "Раим" || acc.Balance != 0 {
		t.Fatalf("NewAccount = %+v, want владельца Раим и нулевой баланс", *acc)
	}
	if err := acc.Deposit(100); err != nil {
		t.Fatalf("Deposit(100): %v", err)
	}
	if err := acc.Deposit(-5); !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("Deposit(-5): err = %v, want ErrInvalidAmount", err)
	}
	if err := acc.Withdraw(0); !errors.Is(err, ErrInvalidAmount) {
		t.Errorf("Withdraw(0): err = %v, want ErrInvalidAmount", err)
	}
	if err := acc.Withdraw(150); !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("Withdraw(150): err = %v, want ErrInsufficientFunds", err)
	}
	if err := acc.Withdraw(30); err != nil {
		t.Fatalf("Withdraw(30): %v", err)
	}
	if acc.Balance != 70 {
		t.Errorf("Balance = %d, want 70", acc.Balance)
	}
}

func TestTransfer(t *testing.T) {
	from, to := NewAccount("А"), NewAccount("Б")
	if from == nil || to == nil {
		t.Fatal("NewAccount вернула nil")
	}
	_ = from.Deposit(50)
	if err := Transfer(from, to, 80); !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("Transfer(80): err = %v, want ErrInsufficientFunds", err)
	}
	if from.Balance != 50 || to.Balance != 0 {
		t.Errorf("неудачный перевод изменил балансы: %d и %d", from.Balance, to.Balance)
	}
	if err := Transfer(from, to, 20); err != nil {
		t.Fatalf("Transfer(20): %v", err)
	}
	if from.Balance != 30 || to.Balance != 20 {
		t.Errorf("после перевода балансы %d и %d, want 30 и 20", from.Balance, to.Balance)
	}
}
