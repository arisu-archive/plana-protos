package protos

import (
	"github.com/arisu-archive/mapx"
)

type EliminateRaidGetBestTeamResponse struct {
	ResponsePacket
	RaidTeamSettingDBsDict *mapx.OrderedMap[string, []RaidTeamSettingDB] `json:",omitempty,omitzero"`
}
