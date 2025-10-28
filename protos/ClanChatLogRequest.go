package protos

import (
	"time"
)

type ClanChatLogRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Channel string `json:",omitempty,omitzero"`
	FromDate time.Time `json:",omitempty,omitzero"`
}
