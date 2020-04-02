package organizer

import "testing"

func TestGuessNumber(t *testing.T) {
	shouldBeOne := guessNumber("random1.xxx")

	if shouldBeOne != "1" {
		t.Errorf("Sum was incorrect, got: %s, want: %s", shouldBeOne, "1")
	}

	guessNumber("random01")

	guessNumber("random")

	guessNumber("random10")
}
