package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type NotificationEventContentReddotResponse struct {
	ResponsePacket
	Reddots                 map[int64][]flatdata.NotificationEventReddot `json:",omitempty,omitzero"`
	EventContentUnlockCGDBs map[int64][]EventContentCollectionDB         `json:",omitempty,omitzero"`
}
