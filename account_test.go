package main

import (
	"testing"
)

func TestInitNewAccount(t *testing.T) {
	u1 := NewAccount("1", 100, "active")
	actual := u1.Balance
	excepted := 100.0
	if actual != excepted {
		t.Errorf("Result was incorrect, got: %f, want %f", actual, excepted)
	}
}

func TestTransferMoney(t *testing.T) {
	u1 := NewAccount("1", 100, "active")
	u2 := NewAccount("2", 0, "active")
	if err := u1.transferMoney("2", 20); err != nil {
		t.Errorf("Should not produce an error")
	}
	excepted := 20.0
	actual := u2.Balance
	if actual != excepted {
		t.Errorf("Result was incorrect, got: %f, want %f", actual, excepted)
	}
}

func TestTransferMoneyFromBlockedAccount(t *testing.T) {
	u1 := NewAccount("1", 100, "blocked")
	u2 := NewAccount("2", 0, "active")
	if err := u1.transferMoney("2", 20); err == nil {
		t.Errorf("Should produce an error")
	}
	if (u1.Balance != 100.0) || (u2.Balance != 0.0) {
		t.Errorf("Balances shouldn't change")
	}
}

func TestTransferMoneyUnknownReceiver(t *testing.T) {
	u1 := NewAccount("1", 100, "active")
	if err := u1.transferMoney("unknown", 20); err == nil {
		t.Errorf("Should produce an error")
	}
}
