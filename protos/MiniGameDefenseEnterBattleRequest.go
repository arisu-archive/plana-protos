package protos

type MiniGameDefenseEnterBattleRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	StageId int64 `json:",omitempty,omitzero"`
}
