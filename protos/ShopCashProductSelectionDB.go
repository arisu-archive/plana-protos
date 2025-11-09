package protos

type ShopCashProductSelectionDB struct {
	SlotGroupId int64 `json:",omitempty,omitzero"`
	SlotIndex   int64 `json:",omitempty,omitzero"`
	SelectionId int64 `json:",omitempty,omitzero"`
}
