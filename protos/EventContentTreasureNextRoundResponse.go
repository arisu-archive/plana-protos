package protos

type EventContentTreasureNextRoundResponse struct {
	ResponsePacket
	BoardHistoryDB EventContentTreasureHistoryDB `json:",omitempty,omitzero"`
	HiddenImage EventContentTreasureCell `json:",omitempty,omitzero"`
}
