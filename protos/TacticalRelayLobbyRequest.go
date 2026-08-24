package protos

type TacticalRelayLobbyRequest struct {
	RequestPacket
	SeasonId int64 `json:",omitempty,omitzero"`
}
