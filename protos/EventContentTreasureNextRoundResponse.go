package protos

type EventContentTreasureNextRoundResponse struct {
	ResponsePacket
	BoardHistoryDB EventContentTreasureHistoryDB
	HiddenImage    EventContentTreasureCell
}
