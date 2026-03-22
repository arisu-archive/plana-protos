package protos

import (
	"github.com/arisu-archive/mapx"
)

type GroupSummary struct {
	TeamId                 int64                                   `json:",omitempty,omitzero"`
	LeaderEntityId         EntityId                                `json:",omitempty,omitzero"`
	Heroes                 []HeroSummary                           `json:",omitempty,omitzero"`
	Supporters             []HeroSummary                           `json:",omitempty,omitzero"`
	CarrierSkillSupporters []HeroSummary                           `json:",omitempty,omitzero"`
	UseAutoSkill           bool                                    `json:",omitempty,omitzero"`
	TSSInteractionServerId int64                                   `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64                                   `json:",omitempty,omitzero"`
	AssistRelations        *mapx.OrderedMap[int64, AssistRelation] `json:",omitempty,omitzero"`
	SkillCostSummary       SkillCostSummary                        `json:",omitempty,omitzero"`
}
