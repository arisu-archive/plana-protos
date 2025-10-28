package protos

import (
	"time"
)

type RequestPacket struct {
	BasePacket
	Resendable bool `json:",omitempty,omitzero"`
	Hash int64 `json:",omitempty,omitzero"`
	IsTest bool `json:",omitempty,omitzero"`
	ModifiedServerTime__DebugOnly *time.Time `json:",omitempty,omitzero"`
}
