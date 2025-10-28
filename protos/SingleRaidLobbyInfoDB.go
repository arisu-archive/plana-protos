package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type SingleRaidLobbyInfoDB struct {
	RaidLobbyInfoDB
	ClearDifficulty []flatdata.Difficulty `json:",omitempty,omitzero"`
}
