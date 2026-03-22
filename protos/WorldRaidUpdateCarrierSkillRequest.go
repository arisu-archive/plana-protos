package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidUpdateCarrierSkillRequest struct {
	RequestPacket
	SeasonId           int64                           `json:",omitempty,omitzero"`
	RecipeIngredientId int64                           `json:",omitempty,omitzero"`
	CarrierSkills      *mapx.OrderedMap[string, int32] `json:",omitempty,omitzero"`
}
