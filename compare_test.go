package matchdata

import (
	"strings"
	"testing"
)

func TestCompareTrackTitles(t *testing.T) {
	testCompareFunc(t, map[string]float64{
		"All You Never Say (Us Version) :: All You Never Say": 100.0,
		"Penthouse Eyes (Lp Version) :: Penthouse Eyes":       100.0,
	}, CompareTrackTitles)
}

func TestCompareAlbumTitles(t *testing.T) {
	testCompareFunc(t, map[string]float64{
		"The Better Life (Deluxe Edition) :: The Better Life":           100.0,
		"Greatest Hits, Vol. 1 (Deluxe Edition) :: Greatest Hits Vol 1": 100.0,
	}, CompareAlbumTitles)
}

func testCompareFunc(t *testing.T, data map[string]float64, compareFunc func(string, string) (float64, float64)) {
	for input, expectedScrubbedScore := range data {
		parts := strings.Split(input, " :: ")
		_, actualScrubbedScore := compareFunc(parts[0], parts[1])
		if expectedScrubbedScore != actualScrubbedScore {
			t.Errorf("Expected '%s' to have score %v. Instead got %v", input, expectedScrubbedScore, actualScrubbedScore)
		}
	}

}
