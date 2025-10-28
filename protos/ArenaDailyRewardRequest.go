package protos

type ArenaDailyRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
