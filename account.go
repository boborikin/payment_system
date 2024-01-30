package main

import (
	"encoding/json"
	"errors"
	"fmt"
)

type account struct {
	Balance float64
	Status  string
}

type transferBody struct {
	Receiver string  `json:"receiver"`
	Value    float64 `json:"value"`
}

const (
	EmissionIBAN    = "BY00CBDC00000000000000000000"
	DestructionIBAN = "BY00CBDC00000000000000000666"
)

const (
	Active  = "active"
	Blocked = "blocked"
)

var (
	NotEnoughMoneyErr          = errors.New("not enough money on account")
	AccountNotActive           = errors.New("account not active")
	AccountReceiverNotFoundErr = errors.New("receiver not found")
	ValueLessOrEqualZero       = errors.New("value less or equal zero")
)

var accounts = make(map[string]*account)

func NewAccount(IBAN string, InitialBalance float64, Status string) *account {
	a := &account{
		Balance: InitialBalance,
		Status:  Status,
	}
	accounts[IBAN] = a
	return a
}

func InitPaymentSystem() {
	NewAccount(EmissionIBAN, 0, Active)
	NewAccount(DestructionIBAN, 0, Active)
}

func (a *account) getEmissionIBAN() string {
	return EmissionIBAN
}

func (a *account) getDestructionIBAN() string {
	return DestructionIBAN
}

func (a *account) destructMoney(value float64) error {
	destructionIBAN := a.getDestructionIBAN()
	if a.Balance >= value {
		a.Balance -= value
		accounts[destructionIBAN].Balance += value
	} else {
		return NotEnoughMoneyErr
	}
	return nil
}

func (a *account) addEmissionBalance(value float64) {
	emissionIBAN := a.getEmissionIBAN()
	accounts[emissionIBAN].Balance += value
}

func (a *account) transferMoney(receiver string, value float64) error {
	// проверка на то, что значение не отрицательное или не равно нулю
	if value <= 0 {
		return ValueLessOrEqualZero
	}
	// проверка на наличие аккаунта получателя
	if _, ok := accounts[receiver]; !ok {
		return AccountReceiverNotFoundErr
	}
	// проверка стутуса на активность
	if a.Status != Active {
		return AccountNotActive
	}
	// проверка на достаточность средств на балансе и перевод средств
	if a.Balance >= value {
		a.Balance -= value
		accounts[receiver].Balance += value
	} else {
		return NotEnoughMoneyErr
	}
	return nil
}

func (a *account) transferMoneyJSON(body string) error {
	tb := transferBody{}
	if err := json.Unmarshal([]byte(body), &tb); err != nil {
		return err
	}
	if err := a.transferMoney(tb.Receiver, tb.Value); err != nil {
		return err
	}
	return nil
}

func AllAccountsInfo() string {
	data, err := json.Marshal(accounts)
	if err != nil {
		fmt.Println(err)
	}
	return string(data)
}
