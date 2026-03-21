package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WorldRaidLocalBossDB struct {
	ContentType  flatdata.ContentType `json:",omitempty,omitzero"`
	SeasonId     int64                `json:",omitempty,omitzero"`
	GroupId      int64                `json:",omitempty,omitzero"`
	UniqueId     int64                `json:",omitempty,omitzero"`
	IsScenario   bool                 `json:",omitempty,omitzero"`
	IsCleardEver bool                 `json:",omitempty,omitzero"`
	TacticMscSum int64                `json:",omitempty,omitzero"`
	RaidBattleDB RaidBattleDB         `json:",omitempty,omitzero"`
	IsContinue   bool                 `json:",omitempty,omitzero"`
}
