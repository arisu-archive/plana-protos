package protos

type EventContentTreasureFlipResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BoardHistoryDB EventContentTreasureHistoryDB `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
}
