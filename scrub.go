package matchdata

import (
	"regexp"
	"strings"
	"unicode"
)

// ScrubArtistName scrubs noise words for artist names, removes punctuation
// and lowercases values
func ScrubArtistName(original string) string {
	var result string
	result = Scrub(original)
	return result
}

// TrackTitleIgnoredPhrases is the array of regular expressions which will
// be stripped from track titles during the scrubbing process
var TrackTitleIgnoredPhrases = []*regexp.Regexp{
	regexp.MustCompile("(?i)remaster(ed)?"),
	regexp.MustCompile("(?i)parental\\s*advisory"),
	regexp.MustCompile("(?i)(radio|deluxe|album|expanded)\\s*(edit|edition|version|release)?"),
}

// ScrubTrackTitle scrubs noise words for song titles, removes punctuation
// and lowercases
func ScrubTrackTitle(original string) string {
	var result string
	result = original
	for _, re := range TrackTitleIgnoredPhrases {
		result = re.ReplaceAllString(result, "")
	}
	result = Scrub(result)
	return result
}

//AlbumTitleIgnoredPhrases is the arraty of regular expressions which will
// be stripped from album titles during the scrubbing process
var AlbumTitleIgnoredPhrases = []*regexp.Regexp{
	regexp.MustCompile("(?i)parental\\s*advisory"),
	regexp.MustCompile("(?i)\\b(\\d{4})?\\s*remaster(ed)?\\s*(edition|version|release)?"),
	regexp.MustCompile("(?i)\\b(super)?\\s*(deluxe|expanded)\\s*(edition|version|release)?"),
}

// ScrubAlbumTitle scrubs noise words for album titles, removes punctuation
// and lowercases
func ScrubAlbumTitle(original string) string {
	var result string
	result = original
	for _, re := range AlbumTitleIgnoredPhrases {
		result = re.ReplaceAllString(result, "")
	}
	result = Scrub(result)
	return result
}

// Scrub removes all punctuation and lowercases the word. If the original
// string was all punctuation, then it is returned to avoid losing vital
// information (there is a band called !!!, for example).
//
func Scrub(original string) string {
	scrubbed := strings.Map(func(r rune) rune {
		switch {
		case unicode.IsPunct(r):
			return -1
		case unicode.IsUpper(r):
			return unicode.ToLower(r)
		default:
			return r
		}
	}, original)
	if scrubbed == "" {
		return original
	}
	return scrubbed
}
