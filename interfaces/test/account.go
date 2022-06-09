package test

type Account struct {
	FirstName string
	LastName  string
}

func (account *Account) ChangeName(firstName, lastName string) {
	account.FirstName = firstName
	account.LastName = lastName
}

func (account Account) String() string {
	return account.FirstName + " - " + account.LastName
}
