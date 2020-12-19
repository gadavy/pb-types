package timestamp

import (
	"encoding/json"
)

func (x *Timestamp) MarshalJSON() ([]byte, error) {
	return json.Marshal(TimeFromProto(x).UTC().Format("2006-01-02T15:04:05Z"))
}
