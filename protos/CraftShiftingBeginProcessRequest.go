package protos

type CraftShiftingBeginProcessRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	SlotId int64 `json:",omitempty,omitzero"`
	RecipeId int64 `json:",omitempty,omitzero"`
	ConsumeRequestDB ConsumeRequestDB `json:",omitempty,omitzero"`
}
