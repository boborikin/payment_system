# Payment System


<details>
<summary>Реализованный функционал:</summary>

- [x] выводить номер специального счета для “эмиссии”;
- [x] выводить номер специального счета для “уничтожения”;
- [x] осуществлять эмиссию, по добавлению на счет “эмиссии” указанной суммы;
- [x] осуществлять отправку определенной суммы денег с указанного счета на счет
“уничтожения”;
- [x] открывать новый счет, вы можете генерировать случайный номер счета или по
вашему алгоритму, или использовать сгенерированный вне вашего класса
номер счета просто как параметр;
- [x] осуществлять перевод заданной суммы денег между двумя указанными
счетами; обеспечить два варианта данной команды:
  - с несколькими параметрами
  - с единственным параметром в формате json (структуру придумайте
       сами);
- [x] выводить список всех счетов, включая специальные, с указанием остатка
     денежных средств на них и их статуса (“активен” или “заблокирован”). Выводить необходимо в формате json.
</details>

### Для запуска кода

```sh
go run .
```

### Для запуска тестов

```shell
go test ./...
```
---

## Документация

---

#### Инициализация системы оплата:

```go
InitPaymentSystem()
```
`при вызове данной функции происходит создание двух основных счетов: "эмиссии" и "уничтожения"`

#### Получение IBAN счета эмиссии:

```go
getEmissionIBAN()
```

#### Получение IBAN счета уничтожения:

```go
getDestructionIBAN()
```

#### Перевод средств на счет "уничтожения":
```go
destructMoney(value float64)
```

#### Добавление средств  на счет "эмиссии":
```go
addEmissionBalance(value float64)
```

#### Перевод средств на другой счет:
```go
transferMoney(receiver string, value float64)
```
`receiver - счет получателя`

#### Второй вариант перевода средств:
```go
transferMoneyJSON(body string)
```
`принимается в себя строку json формата следующего вида:`
```go
`{"receiver": *номер счета*, "value": *количество*}`
```

#### Получение информации обо всех счетах:

```go
ALlAccountsInfo()
```
`Возвращает информацию обо всех счетах, в том числе об специальных; Возвращет значения в формате строки для удобства.`

### Добавлены следующие проверки:

- Значения на отрицательность и ноль при переводе 
- Проверка на наличие аккаунта получателя
- Проверка на достаточность средств на балансе
- Проверка статуса аккаунта (должен быть равен **active**)