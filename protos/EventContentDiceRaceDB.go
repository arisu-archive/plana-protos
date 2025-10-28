package protos

type EventContentDiceRaceDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	Node int64 `json:",omitempty,omitzero"`
	LapCount int64 `json:",omitempty,omitzero"`
	DiceRollCount int64 `json:",omitempty,omitzero"`
	ReceiveRewardLapCount int64 `json:",omitempty,omitzero"`
}
