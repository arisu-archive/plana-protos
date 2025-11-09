package protos

type EventContentDiceRaceRollResponse struct {
	ResponsePacket
	ParcelResultDB            ParcelResultDB             `json:",omitempty,omitzero"`
	DiceRaceDB                EventContentDiceRaceDB     `json:",omitempty,omitzero"`
	DiceResults               []EventContentDiceResult   `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
