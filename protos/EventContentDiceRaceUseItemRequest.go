package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EventContentDiceRaceUseItemRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	DiceRaceResultType flatdata.EventContentDiceRaceResultType `json:",omitempty,omitzero"`
}
