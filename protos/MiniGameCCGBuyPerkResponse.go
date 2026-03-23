package protos

type MiniGameCCGBuyPerkResponse struct {
	ResponsePacket
	Perks                     []int64
	ParcelResultDB            *ParcelResultDB `json:",omitempty,omitzero"`
	EventContentCollectionDBs []*EventContentCollectionDB
}
