package organizer

import "testing"

func TestGuessNumber(t *testing.T) {
	shouldBeOne := guessNumber("random1.xxx")
	if shouldBeOne != "1" {
		t.Errorf("1 incorrect, got: %s, want: %s", shouldBeOne, "1")
	}

	shouldBeOneWithoutZero := guessNumber("random01.xxx")
	if shouldBeOneWithoutZero != "1" {
		t.Errorf("2 incorrect, got: %s, want: %s", shouldBeOneWithoutZero, "1")
	}

	shouldBeEmpty := guessNumber("random.xxx")
	if shouldBeEmpty != "" {
		t.Errorf("3 incorrect, got: %s, want: %s", shouldBeEmpty, "")
	}

	shouldBe999 := guessNumber("random000000000999.xxx")
	if shouldBe999 != "999" {
		t.Errorf("4 incorrect, got: %s, want: %s", shouldBe999, "999")
	}

	withoutDot := guessNumber("random123")
	if withoutDot != "123" {
		t.Errorf("5 incorrect, got: %s, want: %s", withoutDot, "123")
	}
}
