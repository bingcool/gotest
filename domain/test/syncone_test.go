package test

import (
	"fmt"
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}
	once.Do(func() {
		fmt.Println("once1")
		return
	})

	once.Do(func() {
		fmt.Println("once2")
	})
}
