package protos

import (
	"github.com/arisu-archive/mapx"
)

type TimeAttackDungeonLobbyResponse struct {
	ResponsePacket
	RoomDBs                 *mapx.OrderedMap[int64, *TimeAttackDungeonRoomDB]
	PreviousRoomDB          *TimeAttackDungeonRoomDB `json:",omitempty,omitzero"`
	ParcelResultDB          *ParcelResultDB          `json:",omitempty,omitzero"`
	AchieveSeasonBestRecord bool                     `json:",omitempty,omitzero"`
	SeasonBestRecord        int64                    `json:",omitempty,omitzero"`
}
