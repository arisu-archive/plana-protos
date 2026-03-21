package protos

type ClueSearchSaveDB struct {
	EventContentId int64              `json:",omitempty,omitzero"`
	Round          int32              `json:",omitempty,omitzero"`
	SlotDBs        []ClueSearchSlotDB `json:",omitempty,omitzero"`
}
