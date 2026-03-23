package protos

import (
	"github.com/arisu-archive/mapx"
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type NotificationEventContentReddotResponse struct {
	ResponsePacket
	Reddots                 *mapx.OrderedMap[int64, []flatdata.NotificationEventReddot]
	EventContentUnlockCGDBs *mapx.OrderedMap[int64, []*EventContentCollectionDB]
}
