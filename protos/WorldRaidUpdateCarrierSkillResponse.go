package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidUpdateCarrierSkillResponse struct {
	ResponsePacket
	ParcelResultDB ParcelResultDB
	CarrierSkills  *mapx.OrderedMap[string, int32]
}
