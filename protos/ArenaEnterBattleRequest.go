package protos

type ArenaEnterBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	OpponentAccountServerId int64 `json:",omitempty,omitzero"`
	OpponentIndex int64 `json:",omitempty,omitzero"`
}
