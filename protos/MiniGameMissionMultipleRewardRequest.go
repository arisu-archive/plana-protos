package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MiniGameMissionMultipleRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MissionCategory flatdata.MissionCategory `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
}
