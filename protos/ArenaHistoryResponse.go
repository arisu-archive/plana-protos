package protos

type ArenaHistoryResponse struct {
	ResponsePacket
	ArenaHistoryDBs []ArenaHistoryDB `json:",omitempty,omitzero"`
	ArenaDamageReportDB []ArenaDamageReportDB `json:",omitempty,omitzero"`
}
