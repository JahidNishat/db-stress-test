package main

import (
	"fmt"
	"os"
	"sync"
	"time"
)

type Account struct {
	balance int
	id      int
	sync.Mutex
}

func Transfer(from, to *Account, amount int) {
	var first, second *Account
	if from.id < to.id {
		first, second = from, to
	} else {
		first, second = to, from
	}
	first.Lock()
	second.Lock()
	from.balance -= amount
	time.Sleep(1 * time.Millisecond)

	to.balance += amount
	second.Unlock()
	first.Unlock()
}

func main() {
	alice := &Account{balance: 100, id: 1}
	bob := &Account{balance: 100, id: 2}

	go func() {
		for {
			Transfer(alice, bob, 10)
			Transfer(bob, alice, 10)
		}
	}()

	for {
		alice.Lock()
		bob.Lock()
		total := alice.balance + bob.balance
		bob.Unlock()
		alice.Unlock()

		if total != 200 {
			fmt.Printf("FRAUD DETECTED! Total: %d\n", total)
			os.Exit(1)
		}
	}
}
