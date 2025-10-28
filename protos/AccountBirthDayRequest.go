package protos

import (
	"time"
)

type AccountBirthDayRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BirthDay time.Time `json:",omitempty,omitzero"`
}
