package tested_test

import (
	"testing"

	"github.com/fakovacic/test-actions/internal/tested"
)

func TestSum(t *testing.T) {
	total := tested.Sum(5, 5)
	if total != 10 {
		t.Errorf("Sum was incorrect, got: %d, want: %d.", total, 10)
	}
}
