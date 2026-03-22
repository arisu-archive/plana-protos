package protos

type MiniGameCCGBuyPerkResponse struct {
	ResponsePacket
	Perks                     []int64
	ParcelResultDB            ParcelResultDB
	EventContentCollectionDBs []EventContentCollectionDB
}
