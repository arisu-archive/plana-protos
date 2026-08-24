package protos

import (
	"github.com/arisu-archive/mapx"
)

type MissionNotifyUIEnteredResponse struct {
	ResponsePacket
	WelcomeCampaignHistoryDBs *mapx.OrderedMap[int64, []*MissionHistoryDB]
}
