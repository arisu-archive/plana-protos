package protos

type WorldRaidLobbyRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
