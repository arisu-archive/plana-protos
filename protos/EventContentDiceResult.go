package protos

type EventContentDiceResult struct {
	Index      int32        `json:",omitempty,omitzero"`
	MoveAmount int32        `json:",omitempty,omitzero"`
	Rewards    []ParcelInfo `json:",omitempty,omitzero"`
}
