package protos

import (
	"time"
)

type AccountBirthDayRequest struct {
	RequestPacket
	BirthDay time.Time `json:",omitempty,omitzero"`
}
