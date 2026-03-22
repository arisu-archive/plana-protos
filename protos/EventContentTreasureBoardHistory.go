package protos

type EventContentTreasureBoardHistory struct {
	TreasureIds []int64
	NormalCells []EventContentTreasureCell
	Treasures   []EventContentTreasureObject
}
