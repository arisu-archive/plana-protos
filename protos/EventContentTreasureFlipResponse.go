package protos

type EventContentTreasureFlipResponse struct {
	ResponsePacket
	BoardHistoryDB EventContentTreasureHistoryDB
	ParcelResultDB ParcelResultDB
}
