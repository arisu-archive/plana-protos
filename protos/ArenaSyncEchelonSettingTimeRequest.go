package protos

type ArenaSyncEchelonSettingTimeRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
