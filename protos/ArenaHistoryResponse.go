package protos

type ArenaHistoryResponse struct {
	ResponsePacket
	Protocol Protocol `json:",omitempty,omitzero"`
	ArenaHistoryDBs []ArenaHistoryDB `json:",omitempty,omitzero"`
	ArenaDamageReportDB []ArenaDamageReportDB `json:",omitempty,omitzero"`
}
