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

	onlyNumber := guessNumber("54.xx")
	if onlyNumber != "54" {
		t.Errorf("6 incorrect, got: %s, want: %s", onlyNumber, "54")
	}
}
func TestReadAllInside(t *testing.T) {
	//tt()

	allFiles := getAllInStockFolder()
	lenAllFiles := len(allFiles)
	if lenAllFiles != 5 {
		t.Errorf("1 incorrect, got: %d, want: %d", lenAllFiles, 5)
	}

	filtered := filterByDirectory(allFiles)
	lenFiltered := len(filtered)
		if lenFiltered != 3 {
			t.Errorf("2 incorrect, got: %d, want: %d", lenFiltered, 3)
		}
	}

//todelete, just to test the behavior
func TestTt(t *testing.T) {
	readAllInside()
	if 1 != 1 {
		t.Errorf("?")
	}
}
