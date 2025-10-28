package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanCreateRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ClanNickName string `json:",omitempty,omitzero"`
	ClanJoinOption flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
