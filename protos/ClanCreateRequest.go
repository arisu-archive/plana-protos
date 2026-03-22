package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanCreateRequest struct {
	RequestPacket
	ClanNickName   string
	ClanJoinOption flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
