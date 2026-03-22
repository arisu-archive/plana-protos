package protos

import (
	"github.com/arisu-archive/mapx"
)

type AuditGachaStatisticsResponse struct {
	ResponsePacket
	GachaResult *mapx.OrderedMap[int64, int64] `json:",omitempty,omitzero"`
}
