package matchdata

import "testing"

func TestScrubLowercases(t *testing.T) {
	data := map[string]string{
		"P!nk":       "pnk",
		"Elton John": "elton john",
		"bob dylan":  "bob dylan",
		"!!!":        "!!!",
	}
	for src, expected := range data {
		actual := Scrub(src)
		if actual != expected {
			t.Errorf("Expected Scrub('%s') == '%s'. Got '%s' instead.", src, expected, actual)
		}
	}
}

func TestScrubPreservesAllPunctuation(t *testing.T) {
	data := map[string]string{
		"!!!": "!!!",
		"()":  "()",
		"":    "",
	}
	for src, expected := range data {
		actual := Scrub(src)
		if actual != expected {
			t.Errorf("Expected Scrub('%s') == '%s'. Got '%s' instead.", src, expected, actual)
		}
	}
}
