package protos

import (
	"github.com/arisu-archive/mapx"
)

type GachaSimulateCheatResponse struct {
	ResponsePacket
	CharacterIdAndCount *mapx.OrderedMap[int64, int32] `json:",omitempty,omitzero"`
	SimulationCount     int64                          `json:",omitempty,omitzero"`
	GoodsUniqueId       int64                          `json:",omitempty,omitzero"`
	GoodsDevName        string                         `json:",omitempty,omitzero"`
}
