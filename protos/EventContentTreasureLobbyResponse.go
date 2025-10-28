package protos

type EventContentTreasureLobbyResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	BoardHistoryDB EventContentTreasureHistoryDB `json:",omitempty,omitzero"`
	HiddenImage EventContentTreasureCell `json:",omitempty,omitzero"`
	VariationId int64 `json:",omitempty,omitzero"`
}
