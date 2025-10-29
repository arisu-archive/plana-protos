package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EliminateRaidLobbyResponse struct {
	ResponsePacket
	SeasonType flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	RaidGiveUpDB RaidGiveUpDB `json:",omitempty,omitzero"`
	RaidLobbyInfoDB EliminateRaidLobbyInfoDB `json:",omitempty,omitzero"`
	AccountCurrencyDB AccountCurrencyDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
