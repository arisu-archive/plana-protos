package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type ClanSettingRequest struct {
	RequestPacket
	ChangedClanName string
	ChangedNotice   string
	ClanJoinOption  flatdata.ClanJoinOption `json:",omitempty,omitzero"`
}
