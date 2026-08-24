package protos

type MiniGameJankenEndBattleRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	IsClear        bool  `json:",omitempty,omitzero"`
	AutoSelectTurn int32 `json:",omitempty,omitzero"`
	UsedTurn       int32 `json:",omitempty,omitzero"`
	LeftHP         int32 `json:",omitempty,omitzero"`
	Score          int32 `json:",omitempty,omitzero"`
}
