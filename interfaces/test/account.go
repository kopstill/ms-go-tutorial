package test

type Account struct {
	firstName string
	lastName  string
}

func (account *Account) ChangeName(firstName, lastName string) {
	account.firstName = firstName
	account.lastName = lastName
}

func (account Account) String() string {
	return account.firstName + " - " + account.lastName
}
