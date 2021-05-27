package tested_test

import (
	"testing"

	"github.com/fakovacic/test-actions/internal/tested"
)

func TestSum(t *testing.T) {
	t.Parallel()
	t.Run("test", func(t *testing.T) {
		t.Parallel()
		if tested.Sum(5, 5) != 5 {
			t.Errorf("Sum was incorrect,  want: %d.", 10)
		}
	})
}
