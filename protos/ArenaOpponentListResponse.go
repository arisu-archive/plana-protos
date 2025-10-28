package protos

import (
	"time"
)

type ArenaOpponentListResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	PlayerRank int64 `json:",omitempty,omitzero"`
	OpponentUserDBs []ArenaUserDB `json:",omitempty,omitzero"`
	AutoRefreshTime time.Time `json:",omitempty,omitzero"`
}
