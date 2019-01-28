package matchdata

import "github.com/texttheater/golang-levenshtein/levenshtein"

// CompareArtistNames calculates the similarity score
// (100 represents a perfect match) between the two submitted
// artist names. Two scores are returned. The first is the direct
// comparison score. The second is the score after scrubbing
// noise words.
//
func CompareArtistNames(a, b string) (score, scrubbedScore float64) {
	scrubbedA, scrubbedB := ScrubArtistName(a), ScrubArtistName(b)
	score = 1.0 - levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	scrubbedScore = 1.0 - levenshtein.RatioForStrings([]rune(scrubbedA), []rune(scrubbedB), levenshtein.DefaultOptions)
	return score * 100, scrubbedScore * 100
}

// CompareTrackTitles calculates the similarity score
// (100 represents a perfect match) between the two submitted track
// titles. Two scores are returned. The first is the direct comparison,
// the second is the score after scrubbing noise words.
//
func CompareTrackTitles(a, b string) (score, scrubbedScore float64) {
	scrubbedA, scrubbedB := ScrubTrackTitle(a), ScrubTrackTitle(b)
	score = 1.0 - levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	scrubbedScore = 1.0 - levenshtein.RatioForStrings([]rune(scrubbedA), []rune(scrubbedB), levenshtein.DefaultOptions)
	return score * 100, scrubbedScore * 100
}

// CompareAlbumTitles calculates the similarity score
// (100 represents a perfect match) between the two submitted album
// titles. Two scores are returned. The first is the direct comparison,
// the second is the score after scrubbing noise words.
//
func CompareAlbumTitles(a, b string) (score, scrubbedScore float64) {
	scrubbedA, scrubbedB := ScrubAlbumTitle(a), ScrubAlbumTitle(b)
	score = 1.0 - levenshtein.RatioForStrings([]rune(a), []rune(b), levenshtein.DefaultOptions)
	scrubbedScore = 1.0 - levenshtein.RatioForStrings([]rune(scrubbedA), []rune(scrubbedB), levenshtein.DefaultOptions)
	return score * 100, scrubbedScore * 100
}
