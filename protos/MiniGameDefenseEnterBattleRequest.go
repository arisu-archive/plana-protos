package protos

type MiniGameDefenseEnterBattleRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageId        int64 `json:",omitempty,omitzero"`
}
