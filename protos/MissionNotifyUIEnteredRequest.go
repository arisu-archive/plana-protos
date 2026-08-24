package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type MissionNotifyUIEnteredRequest struct {
	RequestPacket
	UIType flatdata.MissionCompleteUIPrefabType `json:",omitempty,omitzero"`
}
