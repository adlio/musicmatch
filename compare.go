package matchdata

import (
	"github.com/gonum/stat"
	"github.com/texttheater/golang-levenshtein/levenshtein"
)

// CompareArtistNames calculates the similarity score (100
// represents a perfect match) between arrays of artist names. Each
// artist in the first slice is compared with each artist in the
// second. The best scores for each artist are kept, and the average
// match of all artists is returned.
func CompareArtistNames(a, b []string) (score, scrubbedScore float64) {

	var bestScore, bestScrubbedScore, thisScore, thisScrubbedScore float64
	scores := make([]float64, len(a))
	scrubbedScores := make([]float64, len(a))

	for i, left := range a {
		bestScore = 0.0
		for _, right := range b {
			thisScore, thisScrubbedScore = CompareArtistName(left, right)
			if thisScrubbedScore > bestScrubbedScore {
				bestScrubbedScore = thisScrubbedScore
				bestScore = thisScore
			}
		}
		scores[i], scrubbedScores[i] = bestScore, bestScrubbedScore
	}

	return stat.Mean(scores, nil), stat.Mean(scrubbedScores, nil)
}

// CompareArtistName calculates the similarity score
// (100 represents a perfect match) between the two submitted
// artist names. Two scores are returned. The first is the direct
// comparison score. The second is the score after scrubbing
// noise words.
//
func CompareArtistName(a, b string) (score, scrubbedScore float64) {
	scrubbedA, scrubbedB := ScrubArtistName(a), ScrubArtistName(b)
	score = levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	scrubbedScore = levenshtein.RatioForStrings([]rune(scrubbedA), []rune(scrubbedB), levenshtein.DefaultOptions)
	return score * 100, scrubbedScore * 100
}

// CompareTrackTitles calculates the similarity score
// (100 represents a perfect match) between the two submitted track
// titles. Two scores are returned. The first is the direct comparison,
// the second is the score after scrubbing noise words.
//
func CompareTrackTitles(a, b string) (score, scrubbedScore float64) {
	scrubbedA, scrubbedB := ScrubTrackTitle(a), ScrubTrackTitle(b)
	score = levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	scrubbedScore = levenshtein.RatioForStrings([]rune(scrubbedA), []rune(scrubbedB), levenshtein.DefaultOptions)
	return score * 100, scrubbedScore * 100
}

// CompareAlbumTitles calculates the similarity score
// (100 represents a perfect match) between the two submitted album
// titles. Two scores are returned. The first is the direct comparison,
// the second is the score after scrubbing noise words.
//
func CompareAlbumTitles(a, b string) (score, scrubbedScore float64) {
	scrubbedA, scrubbedB := ScrubAlbumTitle(a), ScrubAlbumTitle(b)
	score = levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	scrubbedScore = levenshtein.RatioForStrings([]rune(scrubbedA), []rune(scrubbedB), levenshtein.DefaultOptions)
	return score * 100, scrubbedScore * 100
}
