package model

import (
	"math/big"
	"sync"
	"time"

	"github.com/go-shana/samples/go-calculator/internal/config"
)

var (
	globalValue = big.NewInt(0)
	mutex       sync.RWMutex
	lastUpdated time.Time
)

func unsafeLoadValue() *big.Int {
	now := time.Now()

	if now.Sub(lastUpdated) > config.Service.TTL {
		globalValue = big.NewInt(0)
		lastUpdated = now
	}

	return globalValue
}

// Add adds value to global value and returns the result.
func Add(value *big.Int) *big.Int {
	mutex.Lock()
	defer mutex.Unlock()

	old := unsafeLoadValue()
	add := big.NewInt(0)
	add.Add(old, value)
	globalValue = add
	return globalValue
}

// Sub substracts value to global value and returns the result.
func Sub(value *big.Int) *big.Int {
	mutex.Lock()
	defer mutex.Unlock()

	old := unsafeLoadValue()
	sub := big.NewInt(0)
	sub.Sub(old, value)
	globalValue = sub
	return globalValue
}

// Mul multiplies value to global value and returns the result.
func Mul(value *big.Int) *big.Int {
	mutex.Lock()
	defer mutex.Unlock()

	old := unsafeLoadValue()
	mul := big.NewInt(0)
	mul.Mul(old, value)
	globalValue = mul
	return globalValue
}

// Div divides value to global value and returns the result.
func Div(value *big.Int) *big.Int {
	mutex.Lock()
	defer mutex.Unlock()

	old := unsafeLoadValue()
	div := big.NewInt(0)
	div.Div(old, value)
	globalValue = div
	return globalValue
}

// String returns the global value as string.
func String() string {
	mutex.RLock()
	defer mutex.RUnlock()

	value := unsafeLoadValue()
	return value.String()
}
