package protos

type EventContentDiceRaceUseItemResponse struct {
	ResponsePacket
	ParcelResultDB            ParcelResultDB
	DiceRaceDB                EventContentDiceRaceDB
	DiceResults               []EventContentDiceResult
	EventContentCollectionDBs []EventContentCollectionDB
}
