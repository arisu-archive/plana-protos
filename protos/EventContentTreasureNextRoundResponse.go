package protos

type EventContentTreasureNextRoundResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BoardHistoryDB EventContentTreasureHistoryDB `json:",omitempty,omitzero"`
	HiddenImage EventContentTreasureCell `json:",omitempty,omitzero"`
}
