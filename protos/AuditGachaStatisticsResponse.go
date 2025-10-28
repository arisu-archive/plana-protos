package protos

type AuditGachaStatisticsResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	GachaResult map[int64]int64 `json:",omitempty,omitzero"`
}
