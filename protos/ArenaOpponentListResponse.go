package protos

type ArenaOpponentListResponse struct {
	ResponsePacket
	PlayerRank      int64         `json:",omitempty,omitzero"`
	OpponentUserDBs []ArenaUserDB `json:",omitempty,omitzero"`
	AutoRefreshTime MxTime        `json:",omitempty,omitzero"`
}
