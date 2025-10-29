package protos

import (
	"time"
)

type ArenaSyncEchelonSettingTimeResponse struct {
	ResponsePacket
	EchelonSettingTime time.Time `json:",omitempty,omitzero"`
}
