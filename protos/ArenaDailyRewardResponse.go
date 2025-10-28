package protos

import (
	"time"
)

type ArenaDailyRewardResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	DailyRewardActiveTime time.Time `json:",omitempty,omitzero"`
}
