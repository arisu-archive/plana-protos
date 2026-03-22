package protos

import (
	"github.com/arisu-archive/mapx"
)

type ResponsePacket struct {
	BasePacket
	ServerTimeTicks                 int64                                             `json:",omitempty,omitzero"`
	ServerNotification              ServerNotificationFlag                            `json:",omitempty,omitzero"`
	MissionProgressDBs              []MissionProgressDB                               `json:",omitempty,omitzero"`
	EventMissionProgressDBDict      *mapx.OrderedMap[int64, []MissionProgressDB]      `json:",omitempty,omitzero"`
	BattlePassMissionProgressDBDict *mapx.OrderedMap[int64, []MissionProgressDB]      `json:",omitempty,omitzero"`
	StaticOpenConditions            *mapx.OrderedMap[string, OpenConditionLockReason] `json:",omitempty,omitzero"`
}
