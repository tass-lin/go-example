package main

import "fmt"

type Account struct {
	id      string
	name    string
	balance float64
}

func (ac *Account) String() string {
	return fmt.Sprintf("Account(id = %s, name = %s, balance = %.2f)",
		ac.id, ac.name, ac.balance)
}

type CheckingAccount struct {
	Account
	overdraftlimit float64
}

func (ac *CheckingAccount) String() string {
	return fmt.Sprintf("CheckingAccount(id = %s, name = %s, balance = %.2f, overdraftlimit = %.2f)",
		ac.id, ac.name, ac.balance, ac.overdraftlimit)
}

func main() {
	account1 := Account{"1234-5678", "Justin Lin", 1000}
	account2 := CheckingAccount{Account{"1234-5678", "Justin Lin", 1000}, 30000}

	// 顯示 Account(id = 1234-5678, name = Justin Lin, balance = 1000.00)
	fmt.Println(&account1)

	// 顯示 CheckingAccount(id = 1234-5678, name = Justin Lin, balance = 1000.00, overdraftlimit = 30000.00)
	fmt.Println(&account2)
}

//當開發者看到某 API 上定義可以接收某介面型態的值時，應該看看該介面定義了哪些行為，接著看看要傳入的值是否有實作這些行為，這樣就可以了，因為 Go 的介面重點是「行為」，不管 API 上定義的介面型態是什麼，只要行為符合都可以傳入。
//
//也就是說 Go 中，介面型態與行為是分開的，應該重視的只有行為本身，本質上與動態定型語言中只重行為而非型態相同，因此「如何知道某個介面的實現型態有哪些？」、「這個型態實現了哪些介面？」這類問題也就不重要了！