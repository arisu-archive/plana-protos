package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type RaidDB struct {
	Owner RaidMemberDescription `json:",omitempty,omitzero"`
	ContentType flatdata.ContentType `json:",omitempty,omitzero"`
	ServerId int64 `json:",omitempty,omitzero"`
	UniqueId int64 `json:",omitempty,omitzero"`
	SeasonId int64 `json:",omitempty,omitzero"`
	Begin time.Time `json:",omitempty,omitzero"`
	End time.Time `json:",omitempty,omitzero"`
	PlayerCount int64 `json:",omitempty,omitzero"`
	Tags []int32 `json:",omitempty,omitzero"`
	SecretCode string `json:",omitempty,omitzero"`
	RaidState flatdata.RaidStatus `json:",omitempty,omitzero"`
	IsPractice bool `json:",omitempty,omitzero"`
	RaidBossDBs []RaidBossDB `json:",omitempty,omitzero"`
	ParticipateCharacterServerIds map[int64][]int64 `json:",omitempty,omitzero"`
	IsEnterRoom bool `json:",omitempty,omitzero"`
	AccountLevelWhenCreateDB int64 `json:",omitempty,omitzero"`
	ClanAssistUsed bool `json:",omitempty,omitzero"`
}
