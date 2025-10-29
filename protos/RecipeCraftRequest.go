package protos

type RecipeCraftRequest struct {
	RequestPacket
	RecipeCraftUniqueId int64 `json:",omitempty,omitzero"`
	RecipeIngredientUniqueId int64 `json:",omitempty,omitzero"`
}
