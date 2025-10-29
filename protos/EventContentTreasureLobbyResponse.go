package protos

type EventContentTreasureLobbyResponse struct {
	ResponsePacket
	BoardHistoryDB EventContentTreasureHistoryDB `json:",omitempty,omitzero"`
	HiddenImage EventContentTreasureCell `json:",omitempty,omitzero"`
	VariationId int64 `json:",omitempty,omitzero"`
}
