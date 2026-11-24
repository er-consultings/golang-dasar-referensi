package main

import (
	"errors"
	"sync"
	"testing"
)

func TestAccountConcurrentDeposits(t *testing.T) {
	var a Account
	var wg sync.WaitGroup
	for range 500 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			a.Deposit(10)
		}()
	}
	wg.Wait()
	if got := a.Balance(); got != 5000 {
		t.Errorf("Balance() = %d, want 5000", got)
	}
}

func TestAccountNeverGoesNegative(t *testing.T) {
	var a Account
	a.Deposit(100)
	var wg sync.WaitGroup
	var ok Visits
	for range 50 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			if a.Withdraw(10) == nil {
				ok.Inc()
			}
		}()
	}
	wg.Wait()
	if ok.Count() != 10 || a.Balance() != 0 {
		t.Errorf("%d withdrawals succeeded, balance %d; want 10 and 0", ok.Count(), a.Balance())
	}
}

func TestWithdrawTooMuch(t *testing.T) {
	var a Account
	if err := a.Withdraw(1); !errors.Is(err, ErrInsufficientFunds) {
		t.Errorf("Withdraw(1) on an empty account = %v, want ErrInsufficientFunds", err)
	}
}

func TestPriceListConcurrent(t *testing.T) {
	p := NewPriceList()
	var wg sync.WaitGroup
	for i := range 100 {
		wg.Add(2)
		go func() {
			defer wg.Done()
			p.Set("tea", int64(i))
		}()
		go func() {
			defer wg.Done()
			p.Get("tea")
		}()
	}
	wg.Wait()
	if _, ok := p.Get("tea"); !ok {
		t.Error("tea is missing after 100 writes")
	}
}
