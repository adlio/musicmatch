package musicmeta

import (
	"encoding/json"
	"encoding/xml"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestScanNull(t *testing.T) {
	d := ReleaseDate{}
	d.Scan(nil)
	assert.True(t, d.IsZero())
}

func TestScanByteSlice(t *testing.T) {
	var val []byte
	val = []byte("2001-01-01")
	d := ReleaseDate{}
	d.Scan(val)
	assert.Equal(t, 2001, d.Year())
	assert.Equal(t, time.January, d.Month())
	assert.Equal(t, 1, d.Day())
}

func TestDateFromString(t *testing.T) {
	d := DateFromString("2015-04-01")
	assert.NotNil(t, d)
	assert.Equal(t, 2015, d.Year())
	assert.Equal(t, time.April, d.Month())
	assert.Equal(t, 1, d.Day())
}

func TestDateFromStringWithBlankTime(t *testing.T) {
	d := DateFromString("2015-03-07 00:00:00")
	assert.NotNil(t, d)
	assert.Equal(t, 2015, d.Year())
	assert.Equal(t, time.March, d.Month())
	assert.Equal(t, 7, d.Day())
}

func TestDateFromStringBlank(t *testing.T) {
	d := DateFromString("")
	assert.True(t, d.IsZero())
}

func TestDateStringValid(t *testing.T) {
	d := ReleaseDate{time.Date(2011, time.May, 3, 0, 0, 0, 0, time.UTC)}
	str := d.String()
	assert.Equal(t, "2011-05-03", str)
}

func TestDateStringInvalid(t *testing.T) {
	d := ReleaseDate{}
	str := d.String()
	assert.Equal(t, "", str)
}

func TestDateScanBlankString(t *testing.T) {
	var d ReleaseDate
	d.Scan("2014-01-01")
	assert.Equal(t, 2014, d.Year())
	assert.Equal(t, time.January, d.Month())
	assert.Equal(t, 1, d.Day())
}

func TestDateScanTime(t *testing.T) {
	var d ReleaseDate
	t2, _ := time.Parse("2006-01-02 15:04:05", "2015-04-01 23:59:00")
	d.Scan(t2)
	assert.Equal(t, 2015, d.Year())
	assert.Equal(t, time.April, d.Month())
	assert.Equal(t, 1, d.Day())
}

func TestDateScanTimelessUTCString(t *testing.T) {
	var d ReleaseDate
	d.Scan("2015-02-02T00:00:00Z")
	assert.Equal(t, 2015, d.Year())
	assert.Equal(t, time.February, d.Month())
	assert.Equal(t, 2, d.Day())
}

func TestDateScanTimelessUTCStringWithNoZ(t *testing.T) {
	var d ReleaseDate
	d.Scan("2006-03-07T00:00:00")
	assert.Equal(t, 2006, d.Year())
	assert.Equal(t, time.March, d.Month())
	assert.Equal(t, 7, d.Day())
}

func TestUnmarshalXML(t *testing.T) {
	s := struct {
		ReleaseDate ReleaseDate
	}{}
	xmlData := []byte(`<Component><ReleaseDate>2013-05-01</ReleaseDate></Component>`)
	err := xml.Unmarshal(xmlData, &s)
	assert.Nil(t, err)
	assert.Equal(t, 2013, s.ReleaseDate.Year())
	assert.Equal(t, time.May, s.ReleaseDate.Month())
	assert.Equal(t, 1, s.ReleaseDate.Day())
}

func TestMarshalJSON(t *testing.T) {
	s := struct {
		ReleaseDate ReleaseDate
	}{ReleaseDate: DateFromString("2013-05-01")}
	bytes, err := json.Marshal(s)
	assert.Nil(t, err)
	assert.Equal(t, []byte("{\"ReleaseDate\":\"2013-05-01\"}"), bytes)
}

func TestUnmarshalJSON(t *testing.T) {
	d := ReleaseDate{}
	err := d.UnmarshalJSON([]byte("\"1995-04-30\""))
	assert.Nil(t, err)
	assert.Equal(t, "1995-04-30", d.String())
}

func TestValue(t *testing.T) {
	loc, err := time.LoadLocation("Local")
	assert.Nil(t, err)
	d := ReleaseDate{time.Date(2001, 1, 1, 0, 0, 0, 0, loc)}
	v, err := d.Value()
	assert.Nil(t, err)
	v2 := v.(time.Time)
	assert.Equal(t, d.Time, v2)
}

func TestValueWithIsZeroDates(t *testing.T) {
	v, err := ReleaseDate{}.Value()
	assert.Nil(t, v)
	assert.Nil(t, err)
}
