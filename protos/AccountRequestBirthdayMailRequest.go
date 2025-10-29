package protos

import (
	"time"
)

type AccountRequestBirthdayMailRequest struct {
	RequestPacket
	Birthday time.Time `json:",omitempty,omitzero"`
}
