package protos

type ArenaSyncEchelonSettingTimeResponse struct {
	ResponsePacket
	EchelonSettingTime MxTime `json:",omitempty,omitzero"`
}
