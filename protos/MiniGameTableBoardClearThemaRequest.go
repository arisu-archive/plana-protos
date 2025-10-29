package protos

type MiniGameTableBoardClearThemaRequest struct {
	RequestPacket
	EventContentId int64 `json:",omitempty,omitzero"`
	PreserveItemEffectUniqueIds []int64 `json:",omitempty,omitzero"`
}
