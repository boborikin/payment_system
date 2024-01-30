package main

import (
	"fmt"
)

func main() {
	// вместо генерации номера счета, использую заготовленные номера.
	firstSampleIBAN := "BY00CBDC00000000000000000002"
	secondSampleIBAN := "BY00CBDC00000000000000000003"
	// инициализирует платежную систему: создает счет "эмитент" и счет "уничтожения"
	InitPaymentSystem()
	// создание двух аккаунтов
	firstAccount := NewAccount(firstSampleIBAN, 100, Active)
	secondAccount := NewAccount(secondSampleIBAN, 0, Active)
	// пример тела json
	body := `{"receiver": "BY00CBDC00000000000000000003", "value": 10}`
	// перевод средств, передавая json строку
	if err := firstAccount.transferMoneyJSON(body); err != nil {
		fmt.Println(err)
	}
	fmt.Println(firstAccount.Balance, secondAccount.Balance) // 90 10
	// перевод средств функцией с параметрами
	if err := secondAccount.transferMoney("BY00CBDC00000000000000000002", 10); err != nil {
		fmt.Println(err)
	}
	fmt.Println(firstAccount.Balance, secondAccount.Balance) // 100 0

	fmt.Println(AllAccountsInfo())

}
