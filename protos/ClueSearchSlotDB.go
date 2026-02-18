package protos

type ClueSearchSlotDB struct {
	SlotNumber  int32 `json:",omitempty,omitzero"`
	ClueId      int64 `json:",omitempty,omitzero"`
	IsSubmitted bool  `json:",omitempty,omitzero"`
}
