package protos

import (
	"github.com/arisu-archive/mapx"
)

type TimeAttackDungeonLobbyResponse struct {
	ResponsePacket
	RoomDBs                 *mapx.OrderedMap[int64, TimeAttackDungeonRoomDB]
	PreviousRoomDB          TimeAttackDungeonRoomDB
	ParcelResultDB          ParcelResultDB
	AchieveSeasonBestRecord bool  `json:",omitempty,omitzero"`
	SeasonBestRecord        int64 `json:",omitempty,omitzero"`
}
