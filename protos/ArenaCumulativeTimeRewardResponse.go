package protos

import (
	"time"
)

type ArenaCumulativeTimeRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	TimeRewardAmount int64 `json:",omitempty,omitzero"`
	TimeRewardLastUpdateTime time.Time `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
}
