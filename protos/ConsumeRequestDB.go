package protos

type ConsumeRequestDB struct {
	ConsumeItemServerIdAndCounts      map[int64]int64 `json:",omitempty,omitzero"`
	ConsumeEquipmentServerIdAndCounts map[int64]int64 `json:",omitempty,omitzero"`
	ConsumeFurnitureServerIdAndCounts map[int64]int64 `json:",omitempty,omitzero"`
}
