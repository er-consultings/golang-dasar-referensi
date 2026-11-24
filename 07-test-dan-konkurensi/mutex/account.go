package main

import (
	"errors"
	"sync"
	"sync/atomic"
)

// ErrInsufficientFunds is returned when a withdrawal is larger than the balance.
var ErrInsufficientFunds = errors.New("insufficient funds")

// UnsafeCounter has no lock. Many goroutines calling Inc at once lose updates.
type UnsafeCounter struct {
	n int
}

func (c *UnsafeCounter) Inc() { c.n++ } // read, add one, write: three steps, not one

// Account is safe to use from many goroutines. The mutex sits right above the
// fields it protects, and both are unexported so nobody can skip the lock.
type Account struct {
	mu      sync.Mutex
	balance int64
}

// Deposit adds amount to the balance.
func (a *Account) Deposit(amount int64) {
	a.mu.Lock()
	defer a.mu.Unlock()
	a.balance += amount
}

// Withdraw takes amount out, or returns ErrInsufficientFunds. Checking and
// subtracting happen under one lock, so two withdrawals cannot both pass the check.
func (a *Account) Withdraw(amount int64) error {
	a.mu.Lock()
	defer a.mu.Unlock()
	if amount > a.balance {
		return ErrInsufficientFunds
	}
	a.balance -= amount
	return nil
}

// Balance returns the current balance.
func (a *Account) Balance() int64 {
	a.mu.Lock()
	defer a.mu.Unlock()
	return a.balance
}

// PriceList is read far more often than it is written, so it uses RWMutex:
// many readers may hold RLock together, a writer waits for them and holds Lock alone.
type PriceList struct {
	mu     sync.RWMutex
	prices map[string]int64
}

// NewPriceList creates an empty price list.
func NewPriceList() *PriceList {
	return &PriceList{prices: map[string]int64{}}
}

// Set stores the price of an item.
func (p *PriceList) Set(item string, price int64) {
	p.mu.Lock()
	defer p.mu.Unlock()
	p.prices[item] = price
}

// Get returns the price of an item and whether it exists.
func (p *PriceList) Get(item string) (int64, bool) {
	p.mu.RLock()
	defer p.mu.RUnlock()
	v, ok := p.prices[item]
	return v, ok
}

// Visits counts with sync/atomic. For a single number, atomic is simpler than a mutex.
type Visits struct {
	n atomic.Int64
}

func (v *Visits) Inc()         { v.n.Add(1) }
func (v *Visits) Count() int64 { return v.n.Load() }
