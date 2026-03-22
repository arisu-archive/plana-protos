package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type WorldRaidBossListRequest struct {
	RequestPacket
	ContentType              flatdata.ContentType
	SeasonId                 int64 `json:",omitempty,omitzero"`
	RequestOnlyWorldBossData bool  `json:",omitempty,omitzero"`
}
