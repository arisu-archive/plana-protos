package protos

import (
	"time"
)

type ArenaOpponentListResponse struct {
	ResponsePacket
	PlayerRank int64 `json:",omitempty,omitzero"`
	OpponentUserDBs []ArenaUserDB `json:",omitempty,omitzero"`
	AutoRefreshTime time.Time `json:",omitempty,omitzero"`
}
