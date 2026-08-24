package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type RaidDB struct {
	Owner                         *RaidMemberDescription `json:",omitempty,omitzero"`
	ContentType                   flatdata.ContentType
	ServerId                      int64  `json:",omitempty,omitzero"`
	UniqueId                      int64  `json:",omitempty,omitzero"`
	SeasonId                      int64  `json:",omitempty,omitzero"`
	Begin                         MxTime `json:",omitempty,omitzero"`
	End                           MxTime `json:",omitempty,omitzero"`
	PlayerCount                   int64  `json:",omitempty,omitzero"`
	Tags                          []int32
	SecretCode                    string
	RaidState                     flatdata.RaidStatus `json:",omitempty,omitzero"`
	IsPractice                    bool                `json:",omitempty,omitzero"`
	RaidBossDBs                   []*RaidBossDB
	ParticipateCharacterServerIds *mapx.OrderedMap[int64, []int64]
	IsEnterRoom                   bool  `json:",omitempty,omitzero"`
	AccountLevelWhenCreateDB      int64 `json:",omitempty,omitzero"`
	ClanAssistUsed                bool  `json:",omitempty,omitzero"`
}
