package protos

type MiniGameTableBoardUseItemRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EventContentId int64 `json:",omitempty,omitzero"`
	ItemSlotIndex int32 `json:",omitempty,omitzero"`
	UsedItemId int64 `json:",omitempty,omitzero"`
	IsDiscard bool `json:",omitempty,omitzero"`
}
