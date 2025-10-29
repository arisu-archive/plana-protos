package protos

import (
	"time"
)

type ArenaDailyRewardResponse struct {
	ResponsePacket
	ParcelResult ParcelResultDB `json:",omitempty,omitzero"`
	DailyRewardActiveTime time.Time `json:",omitempty,omitzero"`
}
