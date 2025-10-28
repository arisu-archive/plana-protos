package protos

import (
	"time"
)

type AccountRequestBirthdayMailRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Birthday time.Time `json:",omitempty,omitzero"`
}
