package protos

type EventContentEnterTacticRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageUniqueId int64 `json:",omitempty,omitzero"`
	EchelonIndex int64 `json:",omitempty,omitzero"`
	EnemyIndex int64 `json:",omitempty,omitzero"`
}
