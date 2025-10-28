package protos

type ConsumeResultDB struct {
	RemovedItemServerIds []int64 `json:",omitempty,omitzero"`
	RemovedEquipmentServerIds []int64 `json:",omitempty,omitzero"`
	RemovedFurnitureServerIds []int64 `json:",omitempty,omitzero"`
	UsedItemServerIdAndRemainingCounts map[int64]int64 `json:",omitempty,omitzero"`
	UsedEquipmentServerIdAndRemainingCounts map[int64]int64 `json:",omitempty,omitzero"`
	UsedFurnitureServerIdAndRemainingCounts map[int64]int64 `json:",omitempty,omitzero"`
}
