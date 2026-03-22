package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WorldRaidLobbyRequest struct {
	RequestPacket
	ContentType flatdata.ContentType
	SeasonId    int64 `json:",omitempty,omitzero"`
}
