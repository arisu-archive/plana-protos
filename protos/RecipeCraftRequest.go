package protos

type RecipeCraftRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	RecipeCraftUniqueId int64 `json:",omitempty,omitzero"`
	RecipeIngredientUniqueId int64 `json:",omitempty,omitzero"`
}
