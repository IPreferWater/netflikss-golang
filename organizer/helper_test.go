package organizer

import "testing"

func TestGuessNumber(t *testing.T) {
	shouldBeOne := guessNumber("random1.xxx")
	if shouldBeOne != "1" {
		t.Errorf("incorrect, got: %s, want: %s", shouldBeOne, "1")
	}

	shouldBeOneWithoutZero := guessNumber("random01.xxx")
	if shouldBeOneWithoutZero != "1" {
		t.Errorf("incorrect, got: %s, want: %s", shouldBeOneWithoutZero, "1")
	}

	shouldBeEmpty := guessNumber("random.xxx")
	if shouldBeEmpty != "" {
		t.Errorf("incorrect, got: %s, want: %s", shouldBeEmpty, "empty string")
	}

	shouldBe999 := guessNumber("random000000000999")
	if shouldBe999 != "999" {
		t.Errorf("incorrect, got: %s, want: %s", shouldBe999, "999")
	}
}
