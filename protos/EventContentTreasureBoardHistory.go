package protos

type EventContentTreasureBoardHistory struct {
	TreasureIds []int64 `json:",omitempty,omitzero"`
	NormalCells []EventContentTreasureCell `json:",omitempty,omitzero"`
	Treasures []EventContentTreasureObject `json:",omitempty,omitzero"`
}
