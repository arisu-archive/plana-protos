package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanMemberDB struct {
	AccountId                   int64 `json:",omitempty,omitzero"`
	AccountLevel                int64 `json:",omitempty,omitzero"`
	AccountNickName             string
	ClanDBId                    int64                    `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId  int64                    `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64                    `json:",omitempty,omitzero"`
	AttendanceCount             int64                    `json:",omitempty,omitzero"`
	CafeComfortValue            int64                    `json:",omitempty,omitzero"`
	ClanSocialGrade             flatdata.ClanSocialGrade `json:",omitempty,omitzero"`
	JoinDate                    MxTime                   `json:",omitempty,omitzero"`
	SocialGradeUpdateTime       MxTime                   `json:",omitempty,omitzero"`
	LastLoginDate               MxTime                   `json:",omitempty,omitzero"`
	GameLoginDate               MxTime                   `json:",omitempty,omitzero"`
	AppliedDate                 MxTime                   `json:",omitempty,omitzero"`
	AttachmentDB                AccountAttachmentDB
}
