package musicmeta

import (
	"strings"
	"testing"
)

// ScoreDelta is how accurate the score needs to be to
// be considered a match in the below tests
const ScoreDelta = 5.0

func TestCompareArtistNames(t *testing.T) {
	var score, scrubbedScore float64

	score, scrubbedScore = CompareArtistNames(
		[]string{"Rihanna", "Paul McCartney", "Kanye West"},
		[]string{"Kanye West", "Rihanna", "Paul McCartney"},
	)
	if score != 100.0 && scrubbedScore != 100.0 {
		t.Errorf("Expected reordered exact matches to score 100.0. Got %v, %v", score, scrubbedScore)
	}
}

func TestCompareArtistName(t *testing.T) {
	testStringCompareFunc(t, map[string]float64{
		"P!nk :: Pink": 90.0,
	}, CompareArtistName)
}

func TestCompareTrackTitles(t *testing.T) {
	testStringCompareFunc(t, map[string]float64{
		"All You Never Say (Us Version) :: All You Never Say": 100.0,
		"Penthouse Eyes (Lp Version) :: Penthouse Eyes":       100.0,
		"Bean :: Been": 75.0,
	}, CompareTrackTitles)
}

func TestCompareAlbumTitles(t *testing.T) {
	testStringCompareFunc(t, map[string]float64{
		"The Better Life (Deluxe Edition) :: The Better Life":           100.0,
		"Greatest Hits, Vol. 1 (Deluxe Edition) :: Greatest Hits Vol 1": 100.0,
		"Thyme :: Times": 60.0,
	}, CompareAlbumTitles)
}

func TestCompareDurations(t *testing.T) {
	var s1, s2 float64

	s1, s2 = CompareDurations(125, 125)
	if s1 != 100 || s2 != 100 {
		t.Errorf("Expected exactly equal values to get 100.0,100.0. Instead got %v,%v", s1, s2)
	}

	s1, s2 = CompareDurations(200, 100)
	if s1 != 50 || s2 != 50 {
		t.Errorf("Expected half-wrong values to get 50.0,50.0. Instead got %v,%v", s1, s2)
	}
}

func testStringCompareFunc(t *testing.T, data map[string]float64, compareFunc func(string, string) (float64, float64)) {
	for input, expectedScrubbedScore := range data {
		parts := strings.Split(input, " :: ")
		_, actualScrubbedScore := compareFunc(parts[0], parts[1])
		if actualScrubbedScore < (expectedScrubbedScore-ScoreDelta) || actualScrubbedScore > (expectedScrubbedScore+ScoreDelta) {
			t.Errorf("Expected '%s' to have score %v. Instead got %v", input, expectedScrubbedScore, actualScrubbedScore)
		}
	}
}
