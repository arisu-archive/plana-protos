package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidLobbyResponse struct {
	ResponsePacket
	SeasonType        flatdata.RaidSeasonType `json:",omitempty,omitzero"`
	RaidGiveUpDB      RaidGiveUpDB
	RaidLobbyInfoDB   SingleRaidLobbyInfoDB
	AccountCurrencyDB AccountCurrencyDB
	ParcelResultDB    ParcelResultDB
}
