package protos

type EventContentDiceRaceUseItemResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResultDB ParcelResultDB `json:",omitempty,omitzero"`
	DiceRaceDB EventContentDiceRaceDB `json:",omitempty,omitzero"`
	DiceResults []EventContentDiceResult `json:",omitempty,omitzero"`
	EventContentCollectionDBs []EventContentCollectionDB `json:",omitempty,omitzero"`
}
