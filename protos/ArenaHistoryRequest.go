package protos

import (
	"time"
)

type ArenaHistoryRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SearchStartDate time.Time `json:",omitempty,omitzero"`
	Count int32 `json:",omitempty,omitzero"`
}
