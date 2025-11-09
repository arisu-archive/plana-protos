package protos

type CardShopElementDB struct {
	EventContentId    int64 `json:",omitempty,omitzero"`
	SlotNumber        int32 `json:",omitempty,omitzero"`
	CardShopElementId int64 `json:",omitempty,omitzero"`
	SoldOut           bool  `json:",omitempty,omitzero"`
}
