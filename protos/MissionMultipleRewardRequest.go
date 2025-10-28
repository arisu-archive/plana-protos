package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MissionMultipleRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	MissionCategory flatdata.MissionCategory `json:",omitempty,omitzero"`
	GuideMissionSeasonId *int64 `json:",omitempty,omitzero"`
	EventContentId *int64 `json:",omitempty,omitzero"`
}
