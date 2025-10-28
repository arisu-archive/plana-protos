package protos

type EventContentTreasureHistoryDB struct {
	EventContentId int64 `json:",omitempty,omitzero"`
	Round int32 `json:",omitempty,omitzero"`
	Board EventContentTreasureBoardHistory `json:",omitempty,omitzero"`
	IsComplete bool `json:",omitempty,omitzero"`
	HintTreasures []EventContentTreasureObject `json:",omitempty,omitzero"`
}
