package protos

type ArenaCumulativeTimeRewardRequest struct {
	RequestPacket
	Protocol Protocol `json:",omitempty,omitzero"`
}
