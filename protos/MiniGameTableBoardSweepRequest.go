package protos

type MiniGameTableBoardSweepRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	PreserveItemEffectUniqueIds []int64 `json:",omitempty,omitzero"`
}
