package protos

type CraftShiftingBeginProcessRequest struct {
	RequestPacket
	SlotId int64 `json:",omitempty,omitzero"`
	RecipeId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
