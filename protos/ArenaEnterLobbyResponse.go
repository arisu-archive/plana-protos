package protos

type ArenaEnterLobbyResponse struct {
	ResponsePacket
	ArenaPlayerInfoDB *ArenaPlayerInfoDB `json:",omitempty,omitzero"`
	OpponentUserDBs   []ArenaUserDB
	MapId             int64  `json:",omitempty,omitzero"`
	AutoRefreshTime   MxTime `json:",omitempty,omitzero"`
}
