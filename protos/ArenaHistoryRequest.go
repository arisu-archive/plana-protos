package protos

import (
	"time"
)

type ArenaHistoryRequest struct {
	RequestPacket
	SearchStartDate time.Time `json:",omitempty,omitzero"`
	Count int32 `json:",omitempty,omitzero"`
}
