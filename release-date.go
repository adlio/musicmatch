package musicmeta

import (
	"database/sql/driver"
	"encoding/json"
	"encoding/xml"
	"time"
)

// ReleaseDate wraps a time.Time with application-specific XML
// and JSON Marshaling, and implements sql.Valuer
type ReleaseDate struct {
	time.Time
}

// DateFromString creates a Date by parsing the supplied string
func DateFromString(str string) ReleaseDate {
	d := ReleaseDate{}
	d.ScanString(str)
	return d
}

// MarshalJSON implements encoding/json.Marshaler
func (date ReleaseDate) MarshalJSON() ([]byte, error) {
	return json.Marshal(date.String())
}

// UnmarshalXML implements encoding/xml.Unmarshaler
func (date *ReleaseDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	var src string
	d.DecodeElement(&src, &start)
	date.ScanString(src)
	return nil
}

// UnmarshalJSON implements encoding/json.Unmarshaler
func (date *ReleaseDate) UnmarshalJSON(b []byte) (err error) {
	var src string
	if err = json.Unmarshal(b, &src); err == nil {
		date.ScanString(src)
	}
	return err
}

func (date ReleaseDate) String() string {
	if date.IsZero() {
		return ""
	}
	return date.Format("2006-01-02")

}

// Value implements database/sql/driver.Valuer
func (date ReleaseDate) Value() (driver.Value, error) {
	if date.IsZero() {
		return nil, nil
	}
	return date.Time, nil
}

// Scan implements the database/sql.Scanner interface
func (date *ReleaseDate) Scan(value interface{}) error {
	if value == nil {
		return nil
	}

	switch s := value.(type) {
	case []byte:
		date.ScanString(string(s))
	case string:
		date.ScanString(s)
	case time.Time:
		*date = ReleaseDate{s}
	}
	return nil
}

// ScanString is a helper for Scan() to handle string-based input
// parsing to time.Time.
func (date *ReleaseDate) ScanString(s string) {
	var ptn string
	if len(s) == 20 {
		ptn = "2006-01-02T15:04:05Z"
	} else if len(s) == 19 {
		ptn = "2006-01-02T15:04:05"
	} else if len(s) > 10 {
		ptn = "2006-01-02 15:04:05"
	} else {
		ptn = "2006-01-02"
	}
	t, _ := time.ParseInLocation(ptn, s, time.Local)
	*date = ReleaseDate{t}
}
