package account

import "sync"

type Account struct {
	balance      int
	closed       bool
	closeMutex   sync.Mutex
	depositMutex sync.Mutex
}

func Open(amt int) *Account {
	if amt < 0 {
		return nil
	}
	return &Account{
		balance: amt,
		closed:  false,
	}
}

func (ac *Account) Balance() (int, bool) {
	if !ac.closed {
		return ac.balance, true
	}
	return 0, false
}

func (ac *Account) Close() (int, bool) {
	ac.closeMutex.Lock()
	defer ac.closeMutex.Unlock()
	if ac.closed {
		return 0, false
	}
	ac.closed = true
	return ac.balance, true
}

func (ac *Account) Deposit(depAmt int) (int, bool) {
	ac.depositMutex.Lock()
	defer ac.depositMutex.Unlock()
	if ac.closed {
		return 0, false
	}
	if depAmt < 0 && ac.balance+depAmt < 0 {
		return ac.balance, false
	}
	ac.balance += depAmt
	return ac.balance, true
}
