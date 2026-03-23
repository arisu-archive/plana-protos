package protos

type EventContentDiceRaceUseItemResponse struct {
	ResponsePacket
	ParcelResultDB            *ParcelResultDB         `json:",omitempty,omitzero"`
	DiceRaceDB                *EventContentDiceRaceDB `json:",omitempty,omitzero"`
	DiceResults               []EventContentDiceResult
	EventContentCollectionDBs []EventContentCollectionDB
}
