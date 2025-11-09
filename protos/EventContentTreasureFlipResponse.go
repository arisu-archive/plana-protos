package protos

type EventContentTreasureFlipResponse struct {
	ResponsePacket
	BoardHistoryDB EventContentTreasureHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB                `json:",omitempty,omitzero"`
}
