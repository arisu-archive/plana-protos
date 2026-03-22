package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanDB struct {
	ClanDBId                                 int64 `json:",omitempty,omitzero"`
	ClanName                                 string
	ClanChannelName                          string
	ClanPresidentNickName                    string
	ClanPresidentRepresentCharacterUniqueId  int64 `json:",omitempty,omitzero"`
	ClanPresidentRepresentCharacterCostumeId int64 `json:",omitempty,omitzero"`
	ClanNotice                               string
	ClanMemberCount                          int64                   `json:",omitempty,omitzero"`
	ClanJoinOption                           flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
