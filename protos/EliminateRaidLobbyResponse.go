package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type EliminateRaidLobbyResponse struct {
	ResponsePacket
	SeasonType        flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	RaidGiveUpDB      RaidGiveUpDB
	RaidLobbyInfoDB   EliminateRaidLobbyInfoDB
	AccountCurrencyDB AccountCurrencyDB
	ParcelResultDB    ParcelResultDB
}
