package account

import "sync"

// Define the Account type here.

type Account struct {
	acctBalance int64
	closed      bool
	mux         sync.Mutex
}

func Open(amount int64) *Account {
	//panic("Please implement the Open function")
	if amount < 0 {
		return nil
	}
	newAccount := Account{
		acctBalance: amount,
		closed:      false,
		//mux:         &sync.Mutex{},
	}
	return &newAccount
}

func (a *Account) Balance() (int64, bool) {
	//panic("Please implement the Balance function")
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed == true {
		return 0, false
	}
	return a.acctBalance, true
}

func (a *Account) Deposit(amount int64) (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed == true {
		return 0, false
	}
	if (a.acctBalance + amount) < 0 {
		return 0, false
	}
	a.acctBalance += amount
	return a.acctBalance, true
}

func (a *Account) Close() (int64, bool) {
	a.mux.Lock()
	defer a.mux.Unlock()
	if a.closed == true {
		return 0, false
	}
	endingBalance := a.acctBalance
	a.acctBalance = 0
	a.closed = true
	return endingBalance, true
}
