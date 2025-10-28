package protos

import (
	"time"
)

type ArenaSyncEchelonSettingTimeResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	EchelonSettingTime time.Time `json:",omitempty,omitzero"`
}
