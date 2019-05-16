package musicmeta

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDurationFromEmptyString(t *testing.T) {
	d := DurationFromString("")
	assert.EqualValues(t, 0, d)
}

func TestDurationFromStringSecondsOnly(t *testing.T) {
	d := DurationFromString("01")
	assert.Equal(t, 1.0, d.Seconds())
	d = DurationFromString("99")
	assert.Equal(t, 99.0, d.Seconds())
	d = DurationFromString("10")
	assert.Equal(t, 10.0, d.Seconds())
}

func TestDurationFromStringWithTypicalSongLengths(t *testing.T) {
	d := DurationFromString("03:01")
	assert.Equal(t, 3.0, math.Floor(d.Minutes()))
	assert.Equal(t, 181.0, d.Seconds())
	d = DurationFromString("09:09")
	assert.Equal(t, 9.0, math.Floor(d.Minutes()))
	assert.Equal(t, 549.0, d.Seconds())
}

func TestDurationFromStringWithTypicalMovieLengths(t *testing.T) {
	d := DurationFromString("02:15:03")
	assert.Equal(t, 2.0, math.Floor(d.Hours()))
	assert.Equal(t, 135.0, math.Floor(d.Minutes()))
	assert.Equal(t, 8103.0, d.Seconds())
	d = DurationFromString("00:00:01")
	assert.Equal(t, 1.0, d.Seconds())
}

func TestDurationFromStringInSeconds(t *testing.T) {
	assert.Equal(t, 1, DurationFromStringInSeconds("1"))
	assert.Equal(t, 1, DurationFromStringInSeconds("01"))
	assert.Equal(t, 1, DurationFromStringInSeconds("00:00:01"))
	assert.Equal(t, 121, DurationFromStringInSeconds("02:01"))
	assert.Equal(t, 7321, DurationFromStringInSeconds("02:02:01"))
}
