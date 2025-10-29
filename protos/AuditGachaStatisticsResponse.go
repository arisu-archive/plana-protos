package protos

type AuditGachaStatisticsResponse struct {
	ResponsePacket
	GachaResult map[int64]int64 `json:",omitempty,omitzero"`
}
