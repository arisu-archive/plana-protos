package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidListRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RaidBossGroup string `json:",omitempty,omitzero"`
	RaidDifficulty flatdata.Difficulty `json:",omitempty,omitzero"`
	RaidRoomSortOption RaidRoomSortOption `json:",omitempty,omitzero"`
}
