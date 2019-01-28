package musicmatch

import "testing"

func TestScrubArtistNames(t *testing.T) {
	testScrubFunction(t, map[string]string{
		"!!!":  "!!!",
		"P!nk": "pnk",
		"Pink": "pink",
	}, ScrubArtistName)
}

func TestScrubAlbumTitles(t *testing.T) {
	testScrubFunction(t, map[string]string{
		"Lead Us Back: Songs Of Worship [Deluxe Edition]": "lead us back songs of worship",
		"World Peace Is None Of Your Business (Deluxe)":   "world peace is none of your business",
		"Hillbilly Deluxe":                 "hillbilly deluxe",
		"In Da Mix (Super Deluxe Edition)": "in da mix",
	}, ScrubAlbumTitle)
}

func TestScrubTrackTitles(t *testing.T) {
	testScrubFunction(t, map[string]string{
		"Blow Me One Last Kiss":                             "blow me one last kiss",
		"Blow Me (One Last Kiss)":                           "blow me one last kiss",
		"Nowhere Fast (2011 Remastered Version)":            "nowhere fast",
		"Somebody To Love (Live Version - Remastered 2004)": "somebody to love",
		"Dancing Barefoot (Digitally Remastered 1996)":      "dancing barefoot",
		"The Message (Re-Recorded / Remastered Version)":    "the message",
		"Dirty Acres [Deluxe Edition]":                      "dirty acres",
		"Little Darlin’ (Re-Recorded / Remastered)":         "little darlin",
	}, ScrubTrackTitle)
}

func TestScrubLowercases(t *testing.T) {
	testScrubFunction(t, map[string]string{
		"P!nk":       "pnk",
		"Elton John": "elton john",
		"bob dylan":  "bob dylan",
		"!!!":        "!!!",
	}, Scrub)
}

func TestScrubPreservesAllPunctuation(t *testing.T) {
	testScrubFunction(t, map[string]string{
		"!!!": "!!!",
		"()":  "()",
		"":    "",
	}, Scrub)
}

func testScrubFunction(t *testing.T, data map[string]string, scrubFunc func(string) string) {
	for src, expected := range data {
		actual := scrubFunc(src)
		if actual != expected {
			t.Errorf("Expected '%s' => '%s'. Got '%s' instead.", src, expected, actual)
		}
	}
}
