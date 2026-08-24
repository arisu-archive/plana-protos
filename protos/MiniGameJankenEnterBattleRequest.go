package protos

type MiniGameJankenEnterBattleRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	StageId        int64 `json:",omitempty,omitzero"`
	Multiplier     int32 `json:",omitempty,omitzero"`
	IsDoubleUp     bool  `json:",omitempty,omitzero"`
}
