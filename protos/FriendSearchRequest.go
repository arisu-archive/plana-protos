package protos

import (
	"github.com/arisu-archive/plana-flatbuffers/go/flatdata"
)

type FriendSearchRequest struct {
	RequestPacket
	FriendCode  string                           `json:",omitempty,omitzero"`
	LevelOption flatdata.FriendSearchLevelOption `json:",omitempty,omitzero"`
}
