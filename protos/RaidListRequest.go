package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidListRequest struct {
	RequestPacket
	RaidBossGroup      string
	RaidDifficulty     flatdata.Difficulty `json:",omitempty,omitzero"`
	RaidRoomSortOption RaidRoomSortOption  `json:",omitempty,omitzero"`
}
