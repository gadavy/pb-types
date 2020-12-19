package timestamp

import (
	"time"
)

func TimeToProto(t time.Time) *Timestamp {
	if t.IsZero() {
		return nil
	}

	return &Timestamp{
		Seconds: t.Unix(),
		Nanos:   int32(t.Nanosecond()),
	}
}

func TimeFromProto(t *Timestamp) time.Time {
	if t == nil {
		return time.Time{}
	}

	return time.Unix(t.Seconds, int64(t.Nanos))
}
