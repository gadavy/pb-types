package timestamp

import (
	"encoding/json"
	"fmt"
	"strings"
	"time"
)

const (
	secondsInNanos       = 999999999
	maxSecondsInDuration = 315576000000
	maxTimestampSeconds  = 253402300799
	minTimestampSeconds  = -62135596800
)

func (x *Timestamp) MarshalJSON() ([]byte, error) {
	if x.Seconds < minTimestampSeconds || x.Seconds > maxTimestampSeconds {
		return nil, fmt.Errorf("seconds out of range %v", x.Seconds)
	}

	if x.Nanos < 0 || x.Nanos > secondsInNanos {
		return nil, fmt.Errorf("nanos out of range %v", x.Nanos)
	}

	// Uses RFC 3339, where generated output will be Z-normalized and uses 0, 3,
	// 6 or 9 fractional digits.
	t := TimeFromProto(x).UTC()
	s := t.Format("2006-01-02T15:04:05.000000000")
	s = strings.TrimSuffix(s, "000")
	s = strings.TrimSuffix(s, "000")
	s = strings.TrimSuffix(s, ".000")

	b, err := json.Marshal(s + "Z")
	if err != nil {
		return nil, fmt.Errorf("can't marshal time string: %w", err)
	}

	return b, nil
}

func (x *Timestamp) UnmarshalJSON(data []byte) error {
	// Ignore null, like in the main JSON package.
	if string(data) == "null" {
		return nil
	}

	var dest time.Time

	if err := json.Unmarshal(data, &dest); err != nil {
		return fmt.Errorf("can't unmarshal time string: %w", err)
	}

	x.Seconds, x.Nanos = dest.Unix(), int32(dest.Nanosecond())

	if x.Seconds < minTimestampSeconds || x.Seconds > maxTimestampSeconds {
		return fmt.Errorf("value out of range: %v", x.Seconds)
	}

	return nil
}
