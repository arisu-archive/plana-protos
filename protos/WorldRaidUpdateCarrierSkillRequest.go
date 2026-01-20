package protos

type WorldRaidUpdateCarrierSkillRequest struct {
	RequestPacket
	SeasonId           int64            `json:",omitempty,omitzero"`
	RecipeIngredientId int64            `json:",omitempty,omitzero"`
	CarrierSkills      map[string]int32 `json:",omitempty,omitzero"`
}
