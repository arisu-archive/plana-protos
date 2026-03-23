package protos

import (
	"github.com/arisu-archive/mapx"
)

type WorldRaidUpdateCarrierSkillResponse struct {
	ResponsePacket
	ParcelResultDB *ParcelResultDB `json:",omitempty,omitzero"`
	CarrierSkills  *mapx.OrderedMap[string, int32]
}
