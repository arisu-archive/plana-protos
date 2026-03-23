package protos

import (
	"github.com/arisu-archive/mapx"
)

type MomoTalkOutLineResponse struct {
	ResponsePacket
	MomoTalkOutLineDBs   []*MomoTalkOutLineDB
	FavorScheduleRecords *mapx.OrderedMap[int64, []int64]
}
