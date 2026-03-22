package protos

import (
	"github.com/arisu-archive/mapx"
)

type GroupSummary struct {
	TeamId                 int64 `json:",omitempty,omitzero"`
	LeaderEntityId         EntityId
	Heroes                 []HeroSummary
	Supporters             []HeroSummary
	CarrierSkillSupporters []HeroSummary
	UseAutoSkill           bool  `json:",omitempty,omitzero"`
	TSSInteractionServerId int64 `json:",omitempty,omitzero"`
	TSSInteractionUniqueId int64 `json:",omitempty,omitzero"`
	AssistRelations        *mapx.OrderedMap[int64, AssistRelation]
	SkillCostSummary       SkillCostSummary
}
