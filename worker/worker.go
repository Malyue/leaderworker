package worker

import (
	"encoding/json"
	"time"
)

type Type string

var (
	Official  Type = "official"
	Candidate Type = "candidate"
	AllTypes       = []Type{Official, Candidate}
)

func (t Type) String() string { return t }
func (t Type) Valid() bool {
	for _, at := range AllTypes {
		if at == t {
			return true
		}
	}
	return false
}

type Worker interface {
	json.Marshaler
	json.Unmarshaler
	GetID() ID
	GetType() Type
	GetCreatedAt() time.Time
	GetStatus() Status
}
