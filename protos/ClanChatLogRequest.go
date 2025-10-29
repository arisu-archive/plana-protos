package protos

import (
	"time"
)

type ClanChatLogRequest struct {
	RequestPacket
	Channel string `json:",omitempty,omitzero"`
	FromDate time.Time `json:",omitempty,omitzero"`
}
