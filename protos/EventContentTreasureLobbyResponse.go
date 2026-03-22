package protos

type EventContentTreasureLobbyResponse struct {
	ResponsePacket
	BoardHistoryDB EventContentTreasureHistoryDB
	HiddenImage    EventContentTreasureCell
	VariationId    int64 `json:",omitempty,omitzero"`
}
