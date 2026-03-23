package protos

type ArenaOpponentListResponse struct {
	ResponsePacket
	PlayerRank      int64 `json:",omitempty,omitzero"`
	OpponentUserDBs []*ArenaUserDB
	AutoRefreshTime MxTime `json:",omitempty,omitzero"`
}
