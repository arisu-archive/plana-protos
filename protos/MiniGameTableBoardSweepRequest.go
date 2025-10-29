package protos

type MiniGameTableBoardSweepRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	PreserveItemEffectUniqueIds []int64 `json:",omitempty,omitzero"`
}
