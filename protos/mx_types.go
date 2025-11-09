package protos

import (
	"encoding/json"
	"strings"
	"time"
)

// MxTime represents a custom time type for MX protocol.
type MxTime time.Time

// MarshalJSON implements the json.Marshaler interface.
func (mt MxTime) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(mt).Format("2006-01-02T15:04:05"))
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (mt *MxTime) UnmarshalJSON(data []byte) error {
	s := strings.Trim(string(data), "\"")
	parsedTime, err := time.Parse("2006-01-02T15:04:05", s)
	if err != nil {
		return err
	}
	*mt = MxTime(parsedTime)
	return nil
}

// Add adds a duration to the MxTime and returns a new MxTime.
func (mt MxTime) Add(d time.Duration) MxTime {
	return MxTime(time.Time(mt).Add(d))
}

func (mt MxTime) IsZero() bool {
	return time.Time(mt).IsZero()
}

func (mt MxTime) After(u MxTime) bool {
	return time.Time(mt).After(time.Time(u))
}

func (mt MxTime) Before(u MxTime) bool {
	return time.Time(mt).Before(time.Time(u))
}

func (mt MxTime) Compare(u MxTime) int {
	return time.Time(mt).Compare(time.Time(u))
}

func (mt MxTime) Equal(u MxTime) bool {
	return time.Time(mt).Equal(time.Time(u))
}
