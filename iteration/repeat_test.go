package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	t.Run("Repeating 5 times", func(t *testing.T) {
		repeated := Repeat("a", 5)
		expected := "aaaaa"

		if repeated != expected {
			t.Errorf("expected %q but got %q", expected, repeated)
		}
	})
}

// Cool thing about benchmarks in Go is that it automatically determines how
// many times it should run. My first run of this ran 13,696,909 times!
func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}

func ExampleRepeat() {
	repeatCount := 5
	characterToRepeat := "e"

	fmt.Println(Repeat(characterToRepeat, repeatCount))
	// Output: eeeee
}
