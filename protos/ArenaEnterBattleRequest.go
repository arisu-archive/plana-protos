package protos

type ArenaEnterBattleRequest struct {
	RequestPacket
	OpponentAccountServerId int64 `json:",omitempty,omitzero"`
	OpponentIndex int64 `json:",omitempty,omitzero"`
}
