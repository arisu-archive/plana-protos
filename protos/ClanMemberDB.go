package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
	"time"
)

type ClanMemberDB struct {
	AccountId int64 `json:",omitempty,omitzero"`
	AccountLevel int64 `json:",omitempty,omitzero"`
	AccountNickName string `json:",omitempty,omitzero"`
	ClanDBId int64 `json:",omitempty,omitzero"`
	RepresentCharacterUniqueId int64 `json:",omitempty,omitzero"`
	RepresentCharacterCostumeId int64 `json:",omitempty,omitzero"`
	AttendanceCount int64 `json:",omitempty,omitzero"`
	CafeComfortValue int64 `json:",omitempty,omitzero"`
	ClanSocialGrade flatdata.ClanSocialGrade `json:",omitempty,omitzero"`
	JoinDate time.Time `json:",omitempty,omitzero"`
	SocialGradeUpdateTime time.Time `json:",omitempty,omitzero"`
	LastLoginDate time.Time `json:",omitempty,omitzero"`
	GameLoginDate time.Time `json:",omitempty,omitzero"`
	AppliedDate time.Time `json:",omitempty,omitzero"`
	AttachmentDB AccountAttachmentDB `json:",omitempty,omitzero"`
}
