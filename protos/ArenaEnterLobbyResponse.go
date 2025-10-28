package protos

import (
	"time"
)

type ArenaEnterLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaPlayerInfoDB ArenaPlayerInfoDB `json:",omitempty,omitzero"`
	OpponentUserDBs []ArenaUserDB `json:",omitempty,omitzero"`
	MapId int64 `json:",omitempty,omitzero"`
	AutoRefreshTime time.Time `json:",omitempty,omitzero"`
}
