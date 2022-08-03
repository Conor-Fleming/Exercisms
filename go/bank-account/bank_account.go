package account

// Define the Account type here.

type Account struct {
	AcctBalance int64
	Closed      bool
}

func Open(amount int64) *Account {
	//panic("Please implement the Open function")
	if amount <= 0 {
		return nil
	}
	newAccount := Account{
		AcctBalance: amount,
		Closed:      false,
	}
	return &newAccount
}

func (a *Account) Balance() (int64, bool) {
	//panic("Please implement the Balance function")
	return 0.0, false
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	//panic("Please implement the Deposit function")
	return 0.0, false
}

func (a *Account) Close() (int64, bool) {
	//panic("Please implement the Close function")
	return 0.0, false
}
