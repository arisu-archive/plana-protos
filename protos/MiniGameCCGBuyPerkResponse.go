package protos

type MiniGameCCGBuyPerkResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	Perks []int64 `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
