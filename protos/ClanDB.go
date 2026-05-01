package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanDB struct {
	ClanDBId                                 int64                   `json:",omitempty,omitzero"`
	ClanName                                 string                  `json:",omitempty,omitzero"`
	ClanChannelName                          string                  `json:",omitempty,omitzero"`
	ClanPresidentNickName                    string                  `json:",omitempty,omitzero"`
	ClanPresidentRepresentCharacterUniqueId  int64                   `json:",omitempty,omitzero"`
	ClanPresidentRepresentCharacterCostumeId int64                   `json:",omitempty,omitzero"`
	ClanNotice                               string                  `json:",omitempty,omitzero"`
	ClanMemberCount                          int64                   `json:",omitempty,omitzero"`
	ClanJoinOption                           flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
