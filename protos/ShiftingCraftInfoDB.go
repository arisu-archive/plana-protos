package protos

type ShiftingCraftInfoDB struct {
	SlotSequence  int64  `json:",omitempty,omitzero"`
	CraftRecipeId int64  `json:",omitempty,omitzero"`
	CraftAmount   int64  `json:",omitempty,omitzero"`
	StartTime     MxTime `json:",omitempty,omitzero"`
	EndTime       MxTime `json:",omitempty,omitzero"`
}
