package protos

import (
	"github.com/arisu-archive/mapx"
)

type MomoTalkFavorScheduleResponse struct {
	ResponsePacket
	ParcelResultDB       ParcelResultDB
	FavorScheduleRecords *mapx.OrderedMap[int64, []int64]
}
