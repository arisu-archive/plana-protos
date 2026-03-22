package protos

import (
	"github.com/arisu-archive/mapx"
)

type MomoTalkOutLineResponse struct {
	ResponsePacket
	MomoTalkOutLineDBs   []MomoTalkOutLineDB              `json:",omitempty,omitzero"`
	FavorScheduleRecords *mapx.OrderedMap[int64, []int64] `json:",omitempty,omitzero"`
}
