package protos

type ArenaEnterLobbyResponse struct {
	ResponsePacket
	ArenaPlayerInfoDB ArenaPlayerInfoDB `json:",omitempty,omitzero"`
	OpponentUserDBs   []ArenaUserDB     `json:",omitempty,omitzero"`
	MapId             int64             `json:",omitempty,omitzero"`
	AutoRefreshTime   MxTime            `json:",omitempty,omitzero"`
}
