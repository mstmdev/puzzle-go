package model

type Account struct {
	Name     string
	age      int
	password string
}

func NewAccount(name string, age int, password string) Account {
	return Account{
		Name:     name,
		age:      age,
		password: password,
	}
}
