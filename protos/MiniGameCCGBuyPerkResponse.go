package protos

type MiniGameCCGBuyPerkResponse struct {
	ResponsePacket
	Perks                     []int64                    `json:",omitempty,omitzero"`
	ParcelResultDB            ParcelResultDB             `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
