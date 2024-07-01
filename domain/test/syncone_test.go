package test

import (
	"sync"
	"testing"
)

func TestOnce(t *testing.T) {
	once := sync.Once{}
	once.Do(func() {
		t.Log("once1")
		return
	})

	once.Do(func() {
		t.Log("once2")
	})
}
