package protos

import (
	"github.com/arisu-archive/mapx"
)

type ResponsePacket struct {
	BasePacket
	ServerTimeTicks                 int64                  `json:",omitempty,omitzero"`
	ServerNotification              ServerNotificationFlag `json:",omitempty,omitzero"`
	MissionProgressDBs              []*MissionProgressDB
	EventMissionProgressDBDict      *mapx.OrderedMap[int64, []*MissionProgressDB]
	BattlePassMissionProgressDBDict *mapx.OrderedMap[int64, []*MissionProgressDB]
	StaticOpenConditions            *mapx.OrderedMap[string, OpenConditionLockReason]
}
