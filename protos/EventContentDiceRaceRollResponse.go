package protos

type EventContentDiceRaceRollResponse struct {
	ResponsePacket
	ParcelResultDB            ParcelResultDB
	DiceRaceDB                EventContentDiceRaceDB
	DiceResults               []EventContentDiceResult
	EventContentCollectionDBs []EventContentCollectionDB
}
