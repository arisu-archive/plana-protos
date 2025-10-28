package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanSettingRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ChangedClanName string `json:",omitempty,omitzero"`
	ChangedNotice string `json:",omitempty,omitzero"`
	ClanJoinOption flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
