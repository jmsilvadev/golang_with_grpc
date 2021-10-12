package providers

import (
	"context"
	"testing"

	"github.com/jmsilvadev/golangtechtask/pkg/internal/storage"
)

func TestNewMockDB(t *testing.T) {
	exp := &storage.MockDb{}
	got := NewMockDB(context.Background())
	if got != exp {
		t.Errorf("Got and Expected are not equals. Got: %v, Exp: %v", got, exp)
	}
}
