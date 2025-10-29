package protos

type ArenaEnterBattlePart1Request struct {
	RequestPacket
	OpponentAccountServerId int64 `json:",omitempty,omitzero"`
	OpponentRank int64 `json:",omitempty,omitzero"`
	OpponentIndex int32 `json:",omitempty,omitzero"`
}
