package protos

type ArenaEnterLobbyResponse struct {
	ResponsePacket
	ArenaPlayerInfoDB ArenaPlayerInfoDB
	OpponentUserDBs   []ArenaUserDB
	MapId             int64  `json:",omitempty,omitzero"`
	AutoRefreshTime   MxTime `json:",omitempty,omitzero"`
}
